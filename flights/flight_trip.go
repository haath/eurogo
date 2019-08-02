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

func (flightTrip *FlightTrip) String() string {

	var str string
	var previous *FlightLeg

	str = fmt.Sprintf("$%v\t", math.Round(flightTrip.Price))

	tripDuration := flightTrip.GetDurationFormatted()
	str = str + fmt.Sprintf("%s\t", tripDuration)

	for _, tripLeg := range flightTrip.Legs {

		flightLeg := tripLeg.(*FlightLeg)

		if previous != nil {

			layoverMinutes := int(flightLeg.Departs().Sub(previous.Arrives()).Minutes())
			layoverFormatted := shared.FormatDuration(layoverMinutes)

			str = str + fmt.Sprintf("-%s-", layoverFormatted)
		}

		str = str + fmt.Sprintf("%s...%s", flightLeg.From(), flightLeg.To())

		previous = flightLeg
	}

	return str
}
