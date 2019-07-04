package flights

import (
	"testing"
	"time"
)

func TestFlightDuration(t *testing.T) {

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

			departs, _ := time.Parse("2006-01-02T15:04:05-07:00", tt.departs)
			arrives, _ := time.Parse("2006-01-02T15:04:05-07:00", tt.arrives)

			flight := Flight{Departs: departs, Arrives: arrives}

			if flight.GetDurationInHours() != tt.hours {
				t.Fatalf("expected %v, got %v", tt.hours, flight.GetDurationInHours())
			}
			if flight.GetDurationInMinutes() != tt.minutes {
				t.Fatalf("expected %v, got %v", tt.minutes, flight.GetDurationInMinutes())
			}
			if flight.GetDurationFormatted() != tt.formatted {
				t.Fatalf("expected %v, got %v", tt.formatted, flight.GetDurationFormatted())
			}
		})
	}

}
