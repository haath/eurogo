package shared

import (
	"math"
	"time"
)

type TripLeg interface {
	From() string
	To() string
	Departs() time.Time
	Arrives() time.Time
}

// GetDuration returns the duration of the flight.
func GetTripDuration(tripLeg TripLeg) time.Duration {

	return tripLeg.Arrives().Sub(tripLeg.Departs())
}

// GetDurationInHours returns the total duration of the flight in hours, rounded.
func GetTripDurationInHours(tripLeg TripLeg) int {

	return int(math.Round(GetTripDuration(tripLeg).Hours()))
}

// GetDurationInMinutes returns the total duration of the flight in minutes.
func GetTripDurationInMinutes(tripLeg TripLeg) int {

	return int(math.Round(GetTripDuration(tripLeg).Minutes()))
}

// GetDurationFormatted returns the duration of the flight in the "2h 30min" format.
func GetTripDurationFormatted(tripLeg TripLeg) string {

	duration := GetTripDurationInMinutes(tripLeg)
	return FormatDuration(duration)
}
