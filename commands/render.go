package commands

import (
	"encoding/json"
	"eurogo/api"
	"eurogo/flights"
	"eurogo/shared"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

func RenderFlightsTable(flightList []flights.FlightTrip) {

	currencyRate := getCurrencyRate()
	currencySymbol := api.CurrencySymbols[Parameters.Currency]

	t := initWriter()
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:  "Price",
			Align: text.AlignRight,
		},
		{
			Name:  "Length",
			Align: text.AlignRight,
		},
	})

	t.AppendHeader(table.Row{"Airline", "Price", "Length", "Date", "Departure", "Arrival", "Trip", "Stops"})

	for _, flight := range flightList {

		price := math.Round(flight.Price * currencyRate)
		priceFormatted := fmt.Sprintf("%v%s", price, currencySymbol)
		departureDate := flight.Departs().Format("Mon 02/01")
		departureTime := flight.Departs().Format("15:04")
		arrivalDate := flight.Arrives().Format("Mon 02/01")
		arrivalTime := flight.Arrives().Format("15:04")

		if arrivalDate != departureDate {
			arrivalTime = arrivalTime + " (+1)"
		}

		t.AppendRow([]interface{}{
			flight.GetAirline(),
			priceFormatted,
			flight.GetDurationFormatted(),
			departureDate,
			departureTime,
			arrivalTime,
			flight.GetLegSummaryString(),
			flight.GetStops(),
		})
	}

	t.Render()
}

func RenderFlightsMonth(flightList []flights.FlightTrip) {

	currencyRate := getCurrencyRate()
	currencySymbol := api.CurrencySymbols[Parameters.Currency]

	t := initWriter()
	t.Style().Options.SeparateRows = true
	t.SetColumnConfigs([]table.ColumnConfig{
		{Name: "Mon", Align: text.AlignCenter}, {Name: "Tue", Align: text.AlignCenter},
		{Name: "Wed", Align: text.AlignCenter}, {Name: "Thu", Align: text.AlignCenter},
		{Name: "Fri", Align: text.AlignCenter}, {Name: "Sat", Align: text.AlignCenter},
		{Name: "Sun", Align: text.AlignCenter},
	})

	t.AppendHeader(table.Row{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"})

	var tableRows []table.Row
	var currentRow table.Row

	for _, flight := range flightList {

		price := math.Round(flight.Price * currencyRate)
		priceFormatted := fmt.Sprintf("%v%s", price, currencySymbol)

		dayOfMonth := flight.Departs().Day()
		dayOfweek := flight.Departs().Weekday()

		cellFormatted := fmt.Sprintf("%d\n%s", dayOfMonth, priceFormatted)

		currentRow = append(currentRow, cellFormatted)

		if dayOfweek == time.Sunday {

			tableRows = append(tableRows, currentRow)
			currentRow = table.Row{}
		}
	}

	// Fill gaps in rows
	for i := range tableRows {
		for len(tableRows[i]) < 7 {
			if i == 0 {
				tableRows[i] = append(table.Row{""}, tableRows[i]...)
			} else {
				tableRows[i] = append(tableRows[i], "")
			}
		}
	}

	t.AppendRows(tableRows)
	t.Render()
}

func RenderFlightsJSON(flightList []flights.FlightTrip) {

	flightsJSON, err := json.MarshalIndent(flightList, "", "    ")
	shared.ErrorHandler(err)
	log.Println(string(flightsJSON))
}

func RenderAirportsJSON(airports []flights.Airport) {

	airportsJSON, err := json.MarshalIndent(airports, "", "    ")
	shared.ErrorHandler(err)
	log.Println(string(airportsJSON))
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
