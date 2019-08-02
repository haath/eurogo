package skiplagged

import (
	"encoding/json"
	"eurogo/api"
	"eurogo/flights"
	"time"
)

// APIBaseURL is the base URL to the Skiplagged API.
const APIBaseURL string = "https://skiplagged.com/"

// APISearchEndpoint is the endpoint to the flight search.
const APISearchEndpoint string = "/api/search.php"

type skiplaggedProvider struct {
}

// SkiplaggedFlightProvider instantiates a client for the Skiplagged API.
func SkiplaggedFlightsProvider() flights.FlightsProvider {

	return &skiplaggedProvider{}
}

func (this *skiplaggedProvider) SearchFlight(from string, to string, departDate time.Time) ([]flights.FlightTrip, error) {

	request, err := api.NewRequest(APIBaseURL)

	if err != nil {
		return nil, err
	}

	request.Endpoint(APISearchEndpoint)
	request.Set("from", from)
	request.Set("to", to)
	request.Set("depart", departDate.Format("2006-01-02"))

	channel := make(chan api.Response)

	go request.Get(channel)

	response := <-channel

	if response.Error != nil {
		return nil, response.Error
	}

	var skiplaggedResponse skiplaggedSearchResponse

	err = json.Unmarshal([]byte(response.Body), &skiplaggedResponse)

	return skiplaggedResponse.getFlights(), err
}
