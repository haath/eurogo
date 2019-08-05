package commands

import (
	"eurogo/api"
	"eurogo/flights"
	"fmt"
	"math"
	"os"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

func RenderFlightsTable(flightList []*flights.FlightTrip) {

	currencyRate := getCurrencyRate()
	currencySymbol := api.CurrencySymbols[Parameters.Currency]

	t := initWriter()
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name: "Price",
			Align: text.AlignRight,
		}
	})

	t.AppendHeader(table.Row{"Price", "Length", "Date", "Departure", "Arrival", "Trip", "Stops"})

	for _, flight := range flightList {

		price := math.Round(flight.Price * currencyRate)
		departureDate := flight.Departs().Format("Mon 02/01")
		departureTime := flight.Departs().Format("15:04")
		arrivalDate := flight.Arrives().Format("Mon 02/01")
		arrivalTime := flight.Arrives().Format("15:04")

		if arrivalDate != departureDate {
			arrivalTime = arrivalTime + " (+1)"
		}

		t.AppendRow([]interface{}{
			fmt.Sprintf("%v%s", price, currencySymbol),
			flight.GetDurationFormatted(),
			departureDate,
			departureTime,
			arrivalTime,
			flight.GetLegSummaryString(),
			len(flight.Legs) - 1,
		})
	}

	t.Render()
}

func getCurrencyRate() float64 {

	if Parameters.Currency == "USD" {

		return 1
	}

	rates := api.GetCurrencyRates("USD")

	return rates[Parameters.Currency]
}

func initWriter() table.Writer {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	return t
}
