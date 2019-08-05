package commands

import (
	"eurogo/flights"
	"eurogo/flights/skiplagged"
	"eurogo/shared"
	"log"
	"time"
)

type FlightCommand struct {
	SearchFilters
	Args     flightPositionalArgs `positional-args:"1" required:"1"`
	OnDates  []string             `long:"on" description:"Add an exact date to the search range."`
	FromDate string               `long:"from" description:"Set a starting date for the search range."`
	ToDate   string               `long:"to" description:"Set an end date for the search range."`
}

type flightPositionalArgs struct {
	From string `positional-arg-name:"<from>"`
	To   string `positional-arg-name:"<two>"`
}

func (cmd *FlightCommand) Execute(args []string) error {

	dates := cmd.getDates()

	flightList := GetFlightsForDates(cmd.Args.From, cmd.Args.To, dates)

	flightList = cmd.SortAndFilter(flightList)

	if Parameters.JSON {

		RenderFlightsJSON(flightList)
		
	} else {

		RenderFlightsTable(flightList)
	}

	return nil
}

func (cmd *FlightCommand) getDates() []time.Time {

	var dates []time.Time

	// Add exact dates.
	for _, dateString := range cmd.OnDates {

		date := shared.ParseInputDate(dateString)
		dates = append(dates, date)
	}

	// Add range of dates.
	if cmd.FromDate != "" {

		var datesToAdd []time.Time
		fromDate := shared.ParseInputDate(cmd.FromDate)

		if cmd.ToDate != "" {

			toDate := shared.ParseInputDate(cmd.ToDate)
			datesToAdd = shared.GetDatesBetween(fromDate, toDate)
		} else {
			log.Fatal("The --from parameter requires a --to.")
		}

		dates = append(dates, datesToAdd...)
	}

	// Create map to keep unique dates.
	datesMap := make(map[time.Time]struct{})
	for _, date := range dates {

		datesMap[date] = struct{}{}
	}

	// Convert map to slice.
	var uniqueDates []time.Time
	for date := range datesMap {

		uniqueDates = append(uniqueDates, date)
	}
	return uniqueDates
}

func GetFlightsForDates(from string, to string, dates []time.Time) []flights.FlightTrip {

	var flightList []flights.FlightTrip
	var requestChannels []chan []flights.FlightTrip

	provider := skiplagged.SkiplaggedFlightsProvider()

	for _, date := range dates {

		channel := make(chan []flights.FlightTrip)
		requestChannels = append(requestChannels, channel)

		go provider.SearchOneway(from, to, date, channel)
	}

	for _, channel := range requestChannels {

		flights := <-channel

		flightList = append(flightList, flights...)
	}

	return flightList
}
