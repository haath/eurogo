package commands

import (
	"eurogo/flights"
	"sort"
)

type SearchFilters struct {
	Count            uint   `short:"c" long:"count" description:"The maximum number of results to return."`
	MaxStops         []uint `long:"max-stops" description:"The maximum amount of stops in the trip."`
	MaxDurationHours uint   `long:"max-hours" description:"The maximum duration for flight trips in hours."`
}

func (filters *SearchFilters) SortAndFilterOneway(flightList []flights.FlightTrip) []flights.FlightTrip {

	sort.Sort(SortFlights(flightList))

	return filters.filter(flightList)
}

func (filters *SearchFilters) SortAndFilterRoundtrip(roundtripList []flights.FlightRoundtrip) []flights.FlightRoundtrip {

	sort.Sort(SortRoundtrips(roundtripList))

	return filters.filterRoundtrips(roundtripList)
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

func (filters *SearchFilters) filterRoundtrips(roundtripList []flights.FlightRoundtrip) []flights.FlightRoundtrip {

	var filteredRoundtripList []flights.FlightRoundtrip

	for _, roundtrip := range roundtripList {

		if (filters.Count == 0 || len(filteredRoundtripList) < int(filters.Count)) &&
			filters.isValid(roundtrip.Outbound) && filters.isValid(roundtrip.Inbound) {

			filteredRoundtripList = append(filteredRoundtripList, roundtrip)
		}
	}

	return filteredRoundtripList
}

func (filters *SearchFilters) isValid(flight flights.FlightTrip) bool {

	maxStops := 0xFF
	if len(filters.MaxStops) > 0 {
		maxStops = int(filters.MaxStops[len(filters.MaxStops)-1])
	}

	maxDurationHours := 0xFF
	if filters.MaxDurationHours > 0 {
		maxDurationHours = int(filters.MaxDurationHours)
	}

	return flight.GetStops() <= maxStops &&
		flight.GetDurationInHours() <= maxDurationHours
}

type SortFlights []flights.FlightTrip

func (a SortFlights) Len() int           { return len(a) }
func (a SortFlights) Less(i, j int) bool { return a[i].Price < a[j].Price }
func (a SortFlights) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type SortRoundtrips []flights.FlightRoundtrip

func (a SortRoundtrips) Len() int { return len(a) }
func (a SortRoundtrips) Less(i, j int) bool {
	return a[i].GetRoundtripPrice() < a[j].GetRoundtripPrice()
}
func (a SortRoundtrips) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
