package flights

import (
	"time"
)

// FlightsProvider is the interface for objects
type FlightsProvider interface {
	SearchOneway(from string, to string, departDate time.Time, flights chan<- []FlightTrip)
	SearchRoundtrip(from string, to string, departDate time.Time, returnDate time.Time, flights chan<- []FlightRoundtrip)

	SearchOnewaySync(from string, to string, departDate time.Time) []FlightTrip
	SearchRoundtripSync(from string, to string, departDate time.Time, returnDate time.Time) []FlightRoundtrip
}
