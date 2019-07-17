package skiplagged

import (
	"eurogo/flights"
	"time"
)

// APIBaseURL is the base URL to the Skiplagged API.
const APIBaseURL string = "https://skiplagged.com/api/search.php"

type skiplaggedFlightProvider struct {
}

func SkiplaggedFlightProvider() flights.FlightsProvider {

	provider := skiplaggedFlightProvider{}
	return &provider
}

func (skiplagged *skiplaggedFlightProvider) Search(from string, to string, departDate time.Time) []flights.Flight {

	var flights []flights.Flight
	return flights
}
