package commands

import (
	"eurogo/flights"
	"sort"
)

type SearchFilters struct {
	Count    uint   `short:"c" long:"count" description:"The maximum number of results to return."`
	MaxStops []uint `long:"max-stops" description:"The maximum amount of stops in the trip."`
}

func (filters *SearchFilters) SortAndFilter(flightList []flights.FlightTrip) []flights.FlightTrip {

	sort.Sort(SortFlights(flightList))

	return filters.filter(flightList)
}

func (filters *SearchFilters) filter(flightList []flights.FlightTrip) []flights.FlightTrip {

	var filteredFlightList []flights.FlightTrip

	for _, flight := range flightList {

		if (filters.Count == 0 || len(filteredFlightList) < int(filters.Count)) &&
			filters.isValid(flight) {

			filteredFlightList = append(filteredFlightList, flight)
		}
	}

	return filteredFlightList
}

func (filters *SearchFilters) isValid(flight flights.FlightTrip) bool {

	maxStops := 0xFF
	if len(filters.MaxStops) > 0 {
		maxStops = int(filters.MaxStops[len(filters.MaxStops)-1])
	}

	return flight.GetStops() <= maxStops
}

type SortFlights []flights.FlightTrip

func (a SortFlights) Len() int           { return len(a) }
func (a SortFlights) Less(i, j int) bool { return a[i].Price < a[j].Price }
func (a SortFlights) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
