package skiplagged

import (
	"eurogo/flights"
	"eurogo/shared"
	"math"
)

type skiplaggedSearchResponse struct {
	Flights  map[string][]interface{} `json:"flights"`
	Depart   [][]interface{}          `json:"depart"`
	Return   [][]interface{}          `json:"return"`
	Duration float64                  `json:"duration"`
}

func (resp *skiplaggedSearchResponse) getDepartFlights() []*flights.FlightTrip {

	return resp.getFlights(false)
}

func (resp *skiplaggedSearchResponse) getReturnFlights() []*flights.FlightTrip {

	return resp.getFlights(true)
}

func (resp *skiplaggedSearchResponse) getFlights(returning bool) []*flights.FlightTrip {

	var flightList []*flights.FlightTrip

	flightRange := resp.Depart
	if returning {
		flightRange = resp.Return
	}

	for _, flight := range flightRange {

		key := flight[3].(string)
		price := math.Round(flight[0].([]interface{})[0].(float64) / 100)

		flightTrip := flights.FlightTrip{Price: price}

		legs := resp.Flights[key][0].([]interface{})

		for _, leg := range legs {

			legData := leg.([]interface{})

			flightLeg := flights.FlightLeg{}

			flightLeg.FlightNumber = legData[0].(string)
			flightLeg.FromAirport = legData[1].(string)
			flightLeg.ToAirport = legData[3].(string)
			flightLeg.Departure = shared.ParseFlightTime(legData[2].(string))
			flightLeg.Arrival = shared.ParseFlightTime(legData[4].(string))

			flightTrip.AddLeg(&flightLeg)
		}

		flightList = append(flightList, &flightTrip)
	}

	return flightList
}
