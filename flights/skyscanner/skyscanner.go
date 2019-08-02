package skyscanner

import (
	"eurogo/api"
	"eurogo/flights"
	"net/url"
	"encoding/json"
)

// APIBaseURL is the base URL to the Skiplagged API.
const APIBaseURL string = "https://www.skyscanner.net/"

// APIAirportsEndpoint is the endpoint to the airport search.
const APIAirportsEndpoint string = "/g/autosuggest-flights/UK/en-GB/"

type skyscannerProvider struct {
}

func SkyscannerAirportsProvider() flights.AirportsProvider {

	return &skyscannerProvider{}
}

func (skyscanner *skyscannerProvider) SearchAirports(query string) ([]flights.Airport, error) {

	request, err := api.NewRequest(APIBaseURL)

	if err != nil {
		return nil, err
	}

	endpoint := APIAirportsEndpoint + url.QueryEscape(query)

	request.Endpoint(endpoint)

	channel := make(chan api.Response)

	go request.Get(channel)

	response := <-channel

	if response.Error != nil {
		return nil, response.Error
	}

	var skyscannerResponse skyscannerAirportsResponse

	err = json.Unmarshal([]byte(response.Body), &skyscannerResponse)

	return skyscannerResponse.getAirports(), err
}
