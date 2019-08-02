package shared

import (
	"fmt"
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
