package skiplagged

import (
	"encoding/json"
	"eurogo/api"
)

type skiplaggedAirportsResponse struct {
	Hints []struct {
		Code string `json:"value"`
		Name string `json:"name"`
	} `json:"hints"`
}

func GetAirportName(airportCode string, airportName chan<- string) {

	request, err := api.NewRequest(APIBaseURL)

	if err != nil {
		airportName <- ""
		return
	}

	request.Endpoint(APIAirportsEndpoint)
	request.Set("term", airportCode)

	channel := make(chan api.Response)

	go request.Get(channel)

	response := <-channel

	if response.Error != nil {
		airportName <- ""
		return
	}

	var skiplaggedAirportsResponse skiplaggedAirportsResponse

	err = json.Unmarshal([]byte(response.Body), &skiplaggedAirportsResponse)

	if err != nil {
		airportName <- ""
		return
	}

	for _, hint := range skiplaggedAirportsResponse.Hints {

		if hint.Code == airportCode {

			airportName <- hint.Name
			return
		}
	}

	airportName <- ""
}
