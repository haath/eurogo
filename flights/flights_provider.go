package flights

import "time"

// FlightsProvider is the interface for objects
type FlightsProvider interface {
	SearchFlight(from string, to string, departDate time.Time, flights chan<- []*FlightTrip)

	SearchFlightSync(from string, to string, departDate time.Time) []*FlightTrip
}
