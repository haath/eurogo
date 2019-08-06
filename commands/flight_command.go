package commands

import (
	"eurogo/flights"
	"eurogo/flights/skiplagged"
	"eurogo/shared"
	"fmt"
	"log"
	"sort"
	"time"
)

type FlightCommand struct {
	SearchFilters
	Args           flightPositionalArgs `positional-args:"1" required:"1"`
	DepartOnDates  []string             `long:"depart-on" description:"Add an exact date to the search range for the outbound flight."`
	DepartFromDate string               `long:"depart-from" description:"Set a starting date for the search range for the outbound flight."`
	DepartToDate   string               `long:"depart-to" description:"Set an end date for the search range for the outbound flight."`
	ReturnOnDates  []string             `long:"return-on" description:"Add an exact date to the search range for the returning flight."`
	ReturnFromDate string               `long:"return-from" description:"Set a starting date for the search range for the returning flight."`
	ReturnToDate   string               `long:"return-to" description:"Set an end date for the search range for the returning flight."`
	Calendar       bool                 `long:"calendar" description:"Output a lowest-fare calendar for roundtrips."`
}

type flightPositionalArgs struct {
	From string `positional-arg-name:"<from>"`
	To   string `positional-arg-name:"<two>"`
}

func (cmd *FlightCommand) Execute(args []string) error {

	outboundDates := getDateRange(cmd.DepartOnDates, cmd.DepartFromDate, cmd.DepartToDate)
	returnDates := getDateRange(cmd.ReturnOnDates, cmd.ReturnFromDate, cmd.ReturnToDate)

	if len(outboundDates) == 0 {

		return fmt.Errorf("please specify outbound dates")
	}

	if len(returnDates) == 0 {

		cmd.executeOneway(outboundDates)

	} else {

		cmd.executeRoundtrip(outboundDates, returnDates)
	}

	return nil
}

func (cmd *FlightCommand) executeOneway(outboundDates []time.Time) {

	flightList := GetOnewayFlightsForDates(cmd.Args.From, cmd.Args.To, outboundDates)

	flightList = cmd.SortAndFilterOneway(flightList)

	if Parameters.JSON {

		RenderJSON(flightList)

	} else {

		RenderFlightsTable(flightList)
	}
}

func (cmd *FlightCommand) executeRoundtrip(outboundDates, inboundDates []time.Time) {

	flightList := GetRoundtripFlightsForDates(cmd.Args.From, cmd.Args.To, outboundDates, inboundDates)

	flightList = cmd.SortAndFilterRoundtrip(flightList)

	if cmd.Calendar {

		calendar := RoundtripsToCalendarCheapest(flightList)

		if Parameters.JSON {

			RenderJSON(calendar)

		} else {

			RenderRoundtripCalendar(calendar)
		}

	} else {

		if Parameters.JSON {

			RenderJSON(flightList)

		} else {

			RenderRoundtripsTable(flightList)
		}
	}
}

func getDateRange(on []string, from, to string) []time.Time {

	var dates []time.Time

	// Add exact dates.
	for _, dateString := range on {

		date := shared.ParseInputDate(dateString)
		dates = append(dates, date)
	}

	// Add range of dates.
	if from != "" {

		var datesToAdd []time.Time
		fromDate := shared.ParseInputDate(from)

		if to != "" {

			toDate := shared.ParseInputDate(to)
			datesToAdd = shared.GetDatesBetween(fromDate, toDate)
		} else {
			log.Fatal("The --from parameter requires a --to.")
		}

		dates = append(dates, datesToAdd...)
	}

	// Create map to keep unique dates.
	datesMap := make(map[time.Time]struct{})
	for _, date := range dates {

		datesMap[date] = struct{}{}
	}

	// Convert map to slice.
	var uniqueDates []time.Time
	for date := range datesMap {

		uniqueDates = append(uniqueDates, date)
	}
	return uniqueDates
}

func GetOnewayFlightsForDates(from string, to string, dates []time.Time) []flights.FlightTrip {

	var flightList []flights.FlightTrip
	var requestChannels []chan []flights.FlightTrip

	provider := skiplagged.SkiplaggedFlightsProvider()

	for _, date := range dates {

		channel := make(chan []flights.FlightTrip)
		requestChannels = append(requestChannels, channel)

		go provider.SearchOneway(from, to, date, channel)
	}

	for _, channel := range requestChannels {

		flights := <-channel

		flightList = append(flightList, flights...)
	}

	return flightList
}

func GetRoundtripFlightsForDates(from string, to string, outboundDates []time.Time, inboundDates []time.Time) []flights.FlightRoundtrip {

	var roundtripList []flights.FlightRoundtrip
	var requestChannels []chan []flights.FlightRoundtrip

	provider := skiplagged.SkiplaggedFlightsProvider()

	for _, outboundDate := range outboundDates {

		for _, inboundDate := range inboundDates {

			channel := make(chan []flights.FlightRoundtrip)
			requestChannels = append(requestChannels, channel)

			go provider.SearchRoundtrip(from, to, outboundDate, inboundDate, channel)
		}

	}

	for _, channel := range requestChannels {

		roundtrips := <-channel

		roundtripList = append(roundtripList, roundtrips...)
	}

	return roundtripList
}

func RoundtripsToCalendarCheapest(roundtrips []flights.FlightRoundtrip) [][]flights.FlightRoundtrip {

	// Create unique map of outbound-inbound pairs.
	roundtripMap := make(map[string]map[string]flights.FlightRoundtrip)

	for _, roundtrip := range roundtrips {

		outboundDate := roundtrip.Outbound.Departs().Format("2006-01-02")
		inboundDate := roundtrip.Inbound.Departs().Format("2006-01-02")

		_, outboundExists := roundtripMap[outboundDate]

		if !outboundExists {
			roundtripMap[outboundDate] = make(map[string]flights.FlightRoundtrip)
		}

		existingRoundtrip, inboundExists := roundtripMap[outboundDate][inboundDate]

		if !inboundExists || roundtrip.IsBetterThan(&existingRoundtrip) {
			roundtripMap[outboundDate][inboundDate] = roundtrip
		}
	}

	// Group by outbound in new map.
	var groupByOutboundMap [][]flights.FlightRoundtrip
	for outboundDate := range roundtripMap {

		outboundRow := []flights.FlightRoundtrip{}

		for _, roundtrip := range roundtripMap[outboundDate] {

			outboundRow = append(outboundRow, roundtrip)
		}

		sort.Sort(sortInboundRowByDate(outboundRow))

		groupByOutboundMap = append(groupByOutboundMap, outboundRow)
	}

	// Sort them by date.
	sort.Sort(sortOuboundRowsByDate(groupByOutboundMap))

	return groupByOutboundMap
}

type sortOuboundRowsByDate [][]flights.FlightRoundtrip

func (a sortOuboundRowsByDate) Len() int { return len(a) }
func (a sortOuboundRowsByDate) Less(i, j int) bool {
	return a[i][0].Outbound.Departs().Sub(a[j][0].Outbound.Departs()) < 0
}
func (a sortOuboundRowsByDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type sortInboundRowByDate []flights.FlightRoundtrip

func (a sortInboundRowByDate) Len() int { return len(a) }
func (a sortInboundRowByDate) Less(i, j int) bool {
	return a[i].Inbound.Departs().Sub(a[j].Inbound.Departs()) < 0
}
func (a sortInboundRowByDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
