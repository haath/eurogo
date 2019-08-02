package skyscanner

import (
	"encoding/json"
	"eurogo/api"
	"eurogo/flights"
	"eurogo/shared"
	"net/url"
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

func (skyscanner *skyscannerProvider) SearchAirports(query string, airports chan<- []*flights.Airport) {

	request, err := api.NewRequest(APIBaseURL)
	shared.ErrorHandler(err)

	endpoint := APIAirportsEndpoint + url.QueryEscape(query)

	request.Endpoint(endpoint)

	channel := make(chan api.Response)

	go request.Get(channel)

	response := <-channel

	shared.ErrorHandler(response.Error)

	var skyscannerResponse skyscannerAirportsResponse

	err = json.Unmarshal([]byte(response.Body), &skyscannerResponse)

	airports <- skyscannerResponse.getAirports()
}

func (skyscanner *skyscannerProvider) SearchAirportsSync(query string) []*flights.Airport {

	airports := make(chan []*flights.Airport)
	go skyscanner.SearchAirports(query, airports)
	return <-airports
}
