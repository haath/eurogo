package skiplagged

import (
	"encoding/json"
	"eurogo/api"
	"eurogo/flights"
	"time"
)

// APIBaseURL is the base URL to the Skiplagged API.
const APIBaseURL string = "https://skiplagged.com/api/search.php"

// APISearchEndpoint is the endpoint to the flight search.
const APISearchEndpoint string = "/api/search.php"

type skiplaggedFlightProvider struct {
}

// SkiplaggedFlightProvider instantiates a client for the Skiplagged API.
func SkiplaggedFlightProvider() flights.FlightsProvider {

	provider := skiplaggedFlightProvider{}
	return &provider
}

func (this *skiplaggedFlightProvider) Search(from string, to string, departDate time.Time) ([]flights.FlightTrip, error) {

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
