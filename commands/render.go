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

	t.AppendHeader(table.Row{"Price", "Airline", "Length", "Date", "Departure", "Arrival", "Trip", "Stops"})

	for _, flight := range flightList {

		price := math.Round(flight.Price * currencyRate)
		priceFormatted := fmt.Sprintf("%v%s", price, currencySymbol)

		t.AppendRow([]interface{}{
			priceFormatted,
			flight.GetAirline(),
			flight.GetDurationFormatted(),
			flight.DepartureDateFormatted(),
			flight.DepartureTimeFormatted(),
			flight.ArrivalTimeFormatted(),
			flight.GetLegSummaryString(),
			flight.GetStops(),
		})
	}

	t.Render()
}

func RenderRoundtripsTable(flightList []flights.FlightRoundtrip) {

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

	t.AppendHeader(table.Row{"Price", "Airline", "Depart", "Outbound Trip", "Length", "Airline", "Return", "Inbound Trip", "Length"})

	for _, roundtrip := range flightList {

		price := math.Round(roundtrip.GetRoundtripPrice() * currencyRate)
		priceFormatted := fmt.Sprintf("%v%s", price, currencySymbol)

		outbound := roundtrip.Outbound
		inbound := roundtrip.Inbound

		t.AppendRow([]interface{}{
			priceFormatted,
			outbound.GetAirline(),
			outbound.DepartureDateFormatted(),
			outbound.GetLegSummaryStringWithTimes(),
			outbound.GetDurationFormatted(),
			inbound.GetAirline(),
			inbound.DepartureDateFormatted(),
			inbound.GetLegSummaryStringWithTimes(),
			inbound.GetDurationFormatted(),
		})
	}

	t.Render()
}

func RenderFlightsMonth(flightList []flights.FlightTrip) {

	currencyRate := getCurrencyRate()
	currencySymbol := api.CurrencySymbols[Parameters.Currency]

	t := initWriter()
	t.Style().Options.SeparateRows = true

	headerRow := table.Row{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.AppendHeader(headerRow)

	centerAll(&t, headerRow)

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

func RenderRoundtripCalendar(calendar [][]flights.FlightRoundtrip, outboundDates, inboundDates []time.Time) {

	currencyRate := getCurrencyRate()
	currencySymbol := api.CurrencySymbols[Parameters.Currency]

	t := initWriter()
	t.Style().Options.SeparateRows = true
	t.Style().Format.Header = text.FormatDefault

	// Create the header row out of the inbound trip dates.
	headerRow := table.Row{"Outbound\\Inbound"}
	for _, inboundDate := range inboundDates {

		header := inboundDate.Format("Mon 02/01")
		headerRow = append(headerRow, header)
	}

	centerAll(&t, headerRow)

	t.AppendHeader(headerRow)

	rowIndex := 1
	for _, outboundRow := range calendar {

		// First column is the outbound date
		outboundDate := outboundRow[0].Outbound.DepartureDateFormatted()
		for outboundDates[rowIndex-1].Format("Mon 02/01") != outboundDate {

			t.AppendRow(table.Row{outboundDates[rowIndex-1].Format("Mon 02/01")})
			rowIndex++
		}

		row := table.Row{outboundDate}

		colIndex := 1
		for _, roundtrip := range outboundRow {

			inboundDate := roundtrip.Inbound.DepartureDateFormatted()
			for headerRow[colIndex] != inboundDate {
				row = append(row, "")
				colIndex++
			}

			price := math.Round(roundtrip.GetRoundtripPrice() * currencyRate)
			priceFormatted := fmt.Sprintf("%v%s", price, currencySymbol)

			cellFormatted := fmt.Sprintf("%s\n%s", priceFormatted, roundtrip.GetLongestDurationFormatted())

			row = append(row, cellFormatted)

			colIndex++
		}

		t.AppendRow(row)
		rowIndex++
	}

	t.Render()
}

func RenderJSON(flightList interface{}) {

	flightsJSON, err := json.MarshalIndent(flightList, "", "    ")
	shared.ErrorHandler(err)
	log.Println(string(flightsJSON))
}

func getCurrencyRate() float64 {

	if Parameters.Currency == "USD" {

		return 1
	}

	rates := api.GetCurrencyRates("USD")

	return rates[Parameters.Currency]
}

func centerAll(t *table.Writer, headerRow table.Row) {

	var configList []table.ColumnConfig

	for _, col := range headerRow {

		config := table.ColumnConfig{
			Name:   col.(string),
			Align:  text.AlignCenter,
			VAlign: text.VAlignMiddle,
		}

		configList = append(configList, config)
	}

	(*t).SetColumnConfigs(configList)
}

func initWriter() table.Writer {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	return t
}
