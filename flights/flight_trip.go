package flights

import (
	"eurogo/shared"
)

type FlightTrip struct {
	shared.Trip
	Price float64 `json:"price"`
}
