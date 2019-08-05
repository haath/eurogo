package skiplagged

import (
	"encoding/json"
	"eurogo/api"
	"eurogo/flights"
	"eurogo/shared"
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

// SkiplaggedFlightsProvider instantiates a client for the Skiplagged API.
func SkiplaggedFlightsProvider() flights.FlightsProvider {

	return &skiplaggedProvider{}
}

func (this *skiplaggedProvider) SearchOneway(from string, to string, departDate time.Time, flights chan<- []flights.FlightTrip) {

	channel := make(chan api.Response)

	this.search(from, to, departDate.Format("2006-01-02"), "", channel)

	response := <-channel
	shared.ErrorHandler(response.Error)

	var skiplaggedResponse skiplaggedSearchResponse

	err := json.Unmarshal([]byte(response.Body), &skiplaggedResponse)
	shared.ErrorHandler(err)

	flights <- skiplaggedResponse.GetOnewayFlights()
}

func (this *skiplaggedProvider) SearchRoundtrip(from string, to string, departDate time.Time, returnDate time.Time, flights chan<- []flights.FlightRoundtrip) {

	channel := make(chan api.Response)

	this.search(from, to, departDate.Format("2006-01-02"), returnDate.Format("2006-01-02"), channel)

	response := <-channel
	shared.ErrorHandler(response.Error)

	var skiplaggedResponse skiplaggedSearchResponse

	err := json.Unmarshal([]byte(response.Body), &skiplaggedResponse)
	shared.ErrorHandler(err)

	flights <- skiplaggedResponse.GetRoundtripFlights()
}

func (this *skiplaggedProvider) SearchOnewaySync(from string, to string, departDate time.Time) []flights.FlightTrip {

	flights := make(chan []flights.FlightTrip)
	go this.SearchOneway(from, to, departDate, flights)
	return <-flights
}

func (this *skiplaggedProvider) SearchRoundtripSync(from string, to string, departDate time.Time, returnDate time.Time) []flights.FlightRoundtrip {

	flights := make(chan []flights.FlightRoundtrip)
	go this.SearchRoundtrip(from, to, departDate, returnDate, flights)
	return <-flights
}

func (this *skiplaggedProvider) search(from string, to string, departDate string, returnDate string, response chan<- api.Response) {

	request, err := api.NewRequest(APIBaseURL)
	shared.ErrorHandler(err)

	request.Endpoint(APISearchEndpoint)
	request.Set("from", from)
	request.Set("to", to)
	request.Set("depart", departDate)
	request.Set("return", returnDate)
	request.Set("sort", "cost")

	go request.Get(response)
}
