package flights

import (
	"encoding/json"
	"eurogo/shared"
	"math"
)

type FlightRoundtrip struct {
	Outbound FlightTrip `json:"outbound"`
	Inbound  FlightTrip `json:"inbound"`
}

func (flightRoundtrip *FlightRoundtrip) GetDurationInSumMinutes() int {

	return flightRoundtrip.Outbound.GetDurationInMinutes() + flightRoundtrip.Inbound.GetDurationInMinutes()
}

func (flightRoundtrip *FlightRoundtrip) GetLongestDurationInMinutes() int {

	return int(math.Max(float64(flightRoundtrip.Outbound.GetDurationInMinutes()), float64(flightRoundtrip.Inbound.GetDurationInMinutes())))
}

func (flightRoundtrip *FlightRoundtrip) GetLongestDurationFormatted() string {

	return shared.FormatDuration(flightRoundtrip.GetLongestDurationInMinutes())
}

func (flightRoundtrip *FlightRoundtrip) GetRoundtripPrice() float64 {

	if flightRoundtrip.sameAirline() && flightRoundtrip.Outbound.RoundtripPrice > 0 {

		return flightRoundtrip.Outbound.RoundtripPrice
	}
	return flightRoundtrip.Outbound.Price + flightRoundtrip.Inbound.Price
}

func (flightRoundtrip *FlightRoundtrip) IsBetterThan(other *FlightRoundtrip) bool {

	if flightRoundtrip.GetRoundtripPrice() < other.GetRoundtripPrice() {

		return true
	}
	if flightRoundtrip.GetRoundtripPrice() == other.GetRoundtripPrice() &&
		flightRoundtrip.GetDurationInSumMinutes() < other.GetDurationInSumMinutes() {

		return true
	}
	return false
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
