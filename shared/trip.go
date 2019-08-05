package shared

import (
	"fmt"
	"math"
	"time"
)

type Trip struct {
	Legs []TripLeg `json:"legs"`
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

func (this *Trip) DepartureDateFormatted() string {

	return this.Departs().Format("Mon 02/01")
}

func (this *Trip) ArrivalDateFormatted() string {

	return this.Arrives().Format("Mon 02/01")
}

func (this *Trip) DepartureTimeFormatted() string {

	return this.Departs().Format("15:04")
}

func (this *Trip) ArrivalTimeFormatted() string {

	str := this.Arrives().Format("15:04")

	dayOffset := 0

	curDate := this.Departs()
	arriveDate := this.Arrives()

	for curDate.Year() != arriveDate.Year() ||
		curDate.Month() != arriveDate.Month() ||
		curDate.Day() != arriveDate.Day() {

		dayOffset++
		curDate = curDate.AddDate(0, 0, 1)
	}

	if dayOffset > 0 {

		str = str + fmt.Sprintf(" (+%d)", dayOffset)
	}

	return str
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
