package commands

import (
	"eurogo/flights"
	"eurogo/shared"
	"sort"
	"time"
)

type FlightMonthCommand struct {
	SearchFilters
	Args flightMonthPositionalArgs `positional-args:"1" required:"1"`
}

type flightMonthPositionalArgs struct {
	From  string `positional-arg-name:"<from>"`
	To    string `positional-arg-name:"<to>"`
	Month string `positional-arg-name:"<YYYY-MM>"`
}

func (cmd *FlightMonthCommand) Execute(args []string) error {

	dates := cmd.getDates()

	flightList := GetOnewayFlightsForDates(cmd.Args.From, cmd.Args.To, dates)

	flightList = cmd.SortAndFilterOneway(flightList)

	flightList = toDailyCheapest(flightList)

	if Parameters.JSON {

		RenderJSON(flightList)

	} else {

		RenderFlightsMonth(flightList)
	}

	return nil
}

func (cmd *FlightMonthCommand) getDates() []time.Time {

	var dates []time.Time

	currentDate, err := time.Parse("2006-01", cmd.Args.Month)
	shared.ErrorHandler(err)

	month := currentDate.Month()

	for currentDate.Month() == month {

		dates = append(dates, currentDate)

		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return dates
}

func toDailyCheapest(flightList []flights.FlightTrip) []flights.FlightTrip {

	datesFlightsMap := make(map[string]flights.FlightTrip)

	for _, flight := range flightList {

		date := flight.Departs().Format("2006-01-02")

		if existingFlight, exists := datesFlightsMap[date]; !exists || flight.Price < existingFlight.Price {

			datesFlightsMap[date] = flight
		}
	}

	var datesFlightsList []flights.FlightTrip
	for _, flight := range datesFlightsMap {

		datesFlightsList = append(datesFlightsList, flight)
	}

	sort.Sort(SortFlightsByDate(datesFlightsList))

	return datesFlightsList
}

type SortFlightsByDate []flights.FlightTrip

func (a SortFlightsByDate) Len() int           { return len(a) }
func (a SortFlightsByDate) Less(i, j int) bool { return a[i].Departs().Sub(a[j].Departs()) < 0 }
func (a SortFlightsByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
