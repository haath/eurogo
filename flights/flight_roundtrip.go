package flights

type FlightRoundtrip struct {
	Outbound FlightTrip `json:"outbound"`
	Inbound  FlightTrip `json:"inbound"`
}

func (flightRoundtrip *FlightRoundtrip) GetRoundtripPrice() float64 {

	if flightRoundtrip.sameAirline() {

		return flightRoundtrip.Outbound.RoundtripPrice
	}
	return flightRoundtrip.Outbound.Price + flightRoundtrip.Inbound.Price
}

func (flightRoundtrip *FlightRoundtrip) sameAirline() bool {

	return flightRoundtrip.Outbound.GetAirline() == flightRoundtrip.Inbound.GetAirline()
}
