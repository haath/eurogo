package shared

import (
	"math"
	"time"
)

type Trip struct {
	Legs []TripLeg
}

func (this *Trip) AddLeg(leg TripLeg) {

	this.Legs = append(this.Legs, leg)
}

func (this *Trip) Departs() time.Time {

	return this.Legs[0].Departs()
}

func (this *Trip) Arrives() time.Time {

	return this.Legs[len(this.Legs)-1].Arrives()
}

// GetDuration returns the duration of the flight.
func (this *Trip) GetDuration() time.Duration {

	return this.Arrives().Sub(this.Departs())
}

// GetDurationInHours returns the total duration of the flight in hours, rounded.
func (this *Trip) GetDurationInHours() int {

	return int(math.Round(this.GetDuration().Hours()))
}

// GetDurationInMinutes returns the total duration of the flight in minutes.
func (this *Trip) GetDurationInMinutes() int {

	return int(math.Round(this.GetDuration().Minutes()))
}

// GetDurationFormatted returns the duration of the flight in the "2h 30min" format.
func (this *Trip) GetDurationFormatted() string {

	duration := this.GetDurationInMinutes()
	return FormatDuration(duration)
}
