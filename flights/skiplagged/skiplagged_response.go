package skiplagged

import (
	"eurogo/flights"
	"eurogo/shared"
)

type skiplaggedSearchResponse struct {
	Flights  map[string][]interface{} `json:"flights"`
	Depart   [][]interface{}          `json:"depart"`
	Duration float64                  `json:"duration"`
}

func (resp *skiplaggedSearchResponse) getFlights() []flights.FlightTrip {

	var flightList []flights.FlightTrip

	for _, depart := range resp.Depart {

		key := depart[3].(string)
		//price := math.Round(depart[0].([]interface{})[0].(float64) / 100)

		flightTrip := flights.FlightTrip{}

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

		flightList = append(flightList, flightTrip)
	}

	return flightList
}
