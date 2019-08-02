package shared

import (
	"testing"
)

func TestTripDuration(t *testing.T) {

	table := []struct {
		departs   string
		arrives   string
		hours     int
		minutes   int
		formatted string
	}{
		{"2016-06-01T22:59:00-07:00", "2016-06-02T06:07:00-04:00", 4, 248, "4h 8min"},
	}

	for _, tt := range table {
		t.Run(tt.departs, func(t *testing.T) {

			departs := ParseFlightTime(tt.departs)
			arrives := ParseFlightTime(tt.arrives)

			trip := TripLeg{Departs: departs, Arrives: arrives}

			if trip.GetDurationInHours() != tt.hours {
				t.Fatalf("expected %v, got %v", tt.hours, trip.GetDurationInHours())
			}
			if trip.GetDurationInMinutes() != tt.minutes {
				t.Fatalf("expected %v, got %v", tt.minutes, trip.GetDurationInMinutes())
			}
			if trip.GetDurationFormatted() != tt.formatted {
				t.Fatalf("expected %v, got %v", tt.formatted, trip.GetDurationFormatted())
			}
		})
	}

}
