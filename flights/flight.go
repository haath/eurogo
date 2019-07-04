package flights

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// Flight represents a single flight from an origin to a destination.
// It does not represent airfare available for purchase, but a part of
// such a trip, which may consist of one or more flights or other transit.
type Flight struct {
	From    string
	To      string
	Airline string
	Departs time.Time
	Arrives time.Time
}

// GetDuration returns the duration of the flight.
func (flight *Flight) GetDuration() time.Duration {

	return flight.Arrives.Sub(flight.Departs)
}

// GetDurationInHours returns the total duration of the flight in hours, rounded.
func (flight *Flight) GetDurationInHours() int {

	return int(math.Round(flight.GetDuration().Hours()))
}

// GetDurationInMinutes returns the total duration of the flight in minutes.
func (flight *Flight) GetDurationInMinutes() int {

	return int(math.Round(flight.GetDuration().Minutes()))
}

// GetDurationFormatted returns the duration of the flight in the "2h 30min" format.
func (flight *Flight) GetDurationFormatted() string {

	duration := flight.GetDurationInMinutes()
	hours := duration / 60
	minutes := duration % 60

	var sb strings.Builder

	if hours > 0 {
		sb.WriteString(fmt.Sprintf("%dh ", hours))
	}
	if minutes > 0 {
		sb.WriteString(fmt.Sprintf("%dmin", minutes))
	}
	return strings.TrimSpace(sb.String())
}
