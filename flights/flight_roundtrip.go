package flights

import "encoding/json"

type FlightRoundtrip struct {
	Outbound FlightTrip `json:"outbound"`
	Inbound  FlightTrip `json:"inbound"`
}

func (flightRoundtrip *FlightRoundtrip) GetRoundtripPrice() float64 {

	if flightRoundtrip.sameAirline() && flightRoundtrip.Outbound.RoundtripPrice > 0 {

		return flightRoundtrip.Outbound.RoundtripPrice
	}
	return flightRoundtrip.Outbound.Price + flightRoundtrip.Inbound.Price
}

func (flightRoundtrip *FlightRoundtrip) sameAirline() bool {

	return flightRoundtrip.Outbound.GetAirline() == flightRoundtrip.Inbound.GetAirline()
}

func (flightRoundtrip *FlightRoundtrip) MarshalJSON() ([]byte, error) {

	type Alias FlightRoundtrip
	return json.Marshal(&struct {
		*Alias
		RoundtripPrice int `json:"roundtrip_price"`
	}{
		Alias:          (*Alias)(flightRoundtrip),
		RoundtripPrice: int(flightRoundtrip.GetRoundtripPrice()),
	})
}
