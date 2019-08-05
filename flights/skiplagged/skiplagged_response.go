package skiplagged

import (
	"eurogo/flights"
	"eurogo/shared"
	"math"
)

type skiplaggedSearchResponse struct {
	Airlines map[string]string        `json:"airlines"`
	Flights  map[string][]interface{} `json:"flights"`
	Depart   [][]interface{}          `json:"depart"`
	Return   [][]interface{}          `json:"return"`
	Duration float64                  `json:"duration"`
}

func (resp *skiplaggedSearchResponse) GetOnewayFlights() []flights.FlightTrip {

	return resp.getOutboundFlights(false)
}

func (resp *skiplaggedSearchResponse) GetRoundtripFlights() []flights.FlightRoundtrip {

	outbound := resp.getOutboundFlights(true)
	inbound := resp.getInboundFlights()

	var roundtrips []flights.FlightRoundtrip

	for _, outboundFlight := range outbound {

		for _, inboundFlight := range inbound {

			roundtripFlight := flights.FlightRoundtrip{
				Outbound: outboundFlight,
				Inbound:  inboundFlight,
			}

			roundtrips = append(roundtrips, roundtripFlight)
		}
	}

	return roundtrips
}

func (resp *skiplaggedSearchResponse) getOutboundFlights(partOfRoundtrip bool) []flights.FlightTrip {

	return resp.getFlights(false, partOfRoundtrip)
}

func (resp *skiplaggedSearchResponse) getInboundFlights() []flights.FlightTrip {

	return resp.getFlights(true, false)
}

func (resp *skiplaggedSearchResponse) getFlights(inbound bool, partOfRoundtrip bool) []flights.FlightTrip {

	var flightList []flights.FlightTrip

	flightRange := resp.Depart
	if inbound {
		flightRange = resp.Return
	}

	for _, flight := range flightRange {

		key := flight[3].(string)

		flightTrip := flights.FlightTrip{}

		priceArr := flight[0].([]interface{})

		if !inbound && partOfRoundtrip && len(priceArr) > 1 {

			flightTrip.Price = math.Round(priceArr[1].(float64) / 100)
			flightTrip.RoundtripPrice = math.Round(priceArr[0].(float64) / 100)

		} else {

			flightTrip.Price = math.Round(priceArr[0].(float64) / 100)
		}

		legs := resp.Flights[key][0].([]interface{})

		for _, leg := range legs {

			legData := leg.([]interface{})

			flightLeg := flights.FlightLeg{}

			flightLeg.FlightNumber = legData[0].(string)
			flightLeg.Airline = resp.getAirlineName(flightLeg.FlightNumber)
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

func (resp *skiplaggedSearchResponse) getAirlineName(flightNumber string) string {

	airlineCode := flightNumber[0:2]
	return resp.Airlines[airlineCode]
}
