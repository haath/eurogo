package flights

import (
	"eurogo/shared"
	"fmt"
	"math"
)

type FlightTrip struct {
	shared.Trip
	Price float64 `json:"price"`
}

func (flightTrip *FlightTrip) GetLegSummaryString() string {

	var str string
	var previous *FlightLeg

	for _, tripLeg := range flightTrip.Legs {

		flightLeg := tripLeg.(*FlightLeg)

		if previous != nil {

			layoverMinutes := int(flightLeg.Departs().Sub(previous.Arrives()).Minutes())
			layoverFormatted := shared.FormatDuration(layoverMinutes)

			str = str + fmt.Sprintf("-[%s]-", layoverFormatted)
		}

		str = str + fmt.Sprintf("%s...%s", flightLeg.From(), flightLeg.To())

		previous = flightLeg
	}
	return str
}

func (flightTrip *FlightTrip) GetAirline() string {

	return flightTrip.Legs[0].(*FlightLeg).Airline
}

func (flightTrip *FlightTrip) GetStops() int {

	return len(flightTrip.Legs) - 1
}

func (flightTrip *FlightTrip) String() string {

	var str string

	str = fmt.Sprintf("$%v\t", math.Round(flightTrip.Price))

	tripDuration := flightTrip.GetDurationFormatted()
	str = str + fmt.Sprintf("%s\t", tripDuration)

	str = str + flightTrip.GetLegSummaryString()

	return str
}
