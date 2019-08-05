package shared

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func FormatDuration(totalMinutes int) string {

	hours := totalMinutes / 60
	minutes := totalMinutes % 60

	var sb strings.Builder

	if hours > 0 {
		sb.WriteString(fmt.Sprintf("%dh ", hours))
	}
	if minutes > 0 {
		sb.WriteString(fmt.Sprintf("%dmin", minutes))
	}
	return strings.TrimSpace(sb.String())
}

func ParseFlightTime(flightTime string) time.Time {

	t, _ := time.Parse("2006-01-02T15:04:05-07:00", flightTime)
	return t
}

func ParseInputDate(inputDate string) time.Time {

	date, err := time.Parse("2006-01-02", inputDate)
	if err != nil {

		log.Fatal("Input dates should be in the YYYY-MM-DD format.")
	}
	return date
}

func IsWorkday(date time.Time) bool {

	// TODO: Add holidays (https://calendarific.com)
	return date.Weekday() != time.Saturday && date.Weekday() != time.Sunday
}

func GetDatesBetween(from time.Time, to time.Time) []time.Time {

	var dates []time.Time
	current := from

	for to.Sub(current) >= 0 {

		dates = append(dates, current)

		current = current.AddDate(0, 0, 1)
	}

	return dates
}

func TimeToDate(t time.Time) time.Time {

	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}
