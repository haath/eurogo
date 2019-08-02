package flights

import "time"

// FlightLeg represents a single flight from an origin to a destination.
// It does not represent airfare available for purchase, but a part of
// such a trip, which may consist of one or more flights or other transit.
type FlightLeg struct {
	FlightNumber string    `json:"flight_number"`
	Airline      string    `json:"airline"`
	FromAirport  string    `json:"from"`
	ToAirport    string    `json:"to"`
	Departure    time.Time `json:"departs"`
	Arrival      time.Time `json:"arrives"`
}

func (flightLeg *FlightLeg) From() string {
	return flightLeg.FromAirport
}
func (flightLeg *FlightLeg) To() string {
	return flightLeg.ToAirport
}
func (flightLeg *FlightLeg) Departs() time.Time {
	return flightLeg.Departure
}
func (flightLeg *FlightLeg) Arrives() time.Time {
	return flightLeg.Arrival
}
