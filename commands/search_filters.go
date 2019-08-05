package commands

import (
	"eurogo/flights"
	"sort"
)

type SearchFilters struct {
	Count uint `short:"c" long:"count" description:"The maximum number of results to return."`
}

func (filters *SearchFilters) SortAndFilter(flightList []*flights.FlightTrip) []*flights.FlightTrip {

	sort.Sort(SortByPrice(flightList))

	return filters.filter(flightList)
}

func (filters *SearchFilters) filter(flightList []*flights.FlightTrip) []*flights.FlightTrip {

	var filteredFlightList []*flights.FlightTrip

	for _, flight := range flightList {

		if len(filteredFlightList) < int(filters.Count) && filters.isValid(flight) {

			filteredFlightList = append(filteredFlightList, flight)
		}
	}

	return filteredFlightList
}

func (filters *SearchFilters) isValid(flightList *flights.FlightTrip) bool {

	return true
}

type SortByPrice []*flights.FlightTrip

func (a SortByPrice) Len() int           { return len(a) }
func (a SortByPrice) Less(i, j int) bool { return a[i].Price < a[j].Price }
func (a SortByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
