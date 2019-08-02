package skiplagged

import (
	"encoding/json"
	"eurogo/shared"
	"eurogo/api"
	"eurogo/flights"
	"time"
)

// APIBaseURL is the base URL to the Skiplagged API.
const APIBaseURL string = "https://skiplagged.com/"

// APISearchEndpoint is the endpoint to the flight search.
const APISearchEndpoint string = "/api/search.php"

// APIAirportsEndpoint is the endpoint to the airport search.
const APIAirportsEndpoint string = "/api/hint.php"

type skiplaggedProvider struct {
}

// SkiplaggedFlightProvider instantiates a client for the Skiplagged API.
func SkiplaggedFlightsProvider() flights.FlightsProvider {

	return &skiplaggedProvider{}
}

func (this *skiplaggedProvider) SearchFlight(from string, to string, departDate time.Time, flights chan<- []*flights.FlightTrip) {

	request, err := api.NewRequest(APIBaseURL)
	shared.ErrorHandler(err)

	request.Endpoint(APISearchEndpoint)
	request.Set("from", from)
	request.Set("to", to)
	request.Set("depart", departDate.Format("2006-01-02"))

	channel := make(chan api.Response)

	go request.Get(channel)

	response := <-channel
	shared.ErrorHandler(response.Error)

	var skiplaggedResponse skiplaggedSearchResponse

	err = json.Unmarshal([]byte(response.Body), &skiplaggedResponse)

	flights <- skiplaggedResponse.getFlights()
}

func (this *skiplaggedProvider) SearchFlightSync(from string, to string, departDate time.Time) []*flights.FlightTrip {

	flights := make(chan []*flights.FlightTrip)
	go this.SearchFlight(from, to, departDate, flights)
	return <-flights
}
