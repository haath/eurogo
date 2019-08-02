package main

import (
	"log"
	"os"
	"time"

	"eurogo/commands"
	"eurogo/flights"
	"eurogo/flights/skiplagged"

	"github.com/jessevdk/go-flags"
)

var Options struct {
	// Commands
	Airport commands.AirportCommand `command:"airport" description:"Search for an airport."`
}

func foo() {

	prov := skiplagged.SkiplaggedFlightsProvider()

	flightsList, err := prov.SearchFlight("PRG", "ATH", time.Date(2019, time.October, 12, 0, 0, 0, 0, time.UTC))

	if err != nil {
		log.Fatal(err)
	}

	flight := flightsList[0]
	log.Println(flight.Price)
	for _, leg := range flight.Legs {

		flightLeg := leg.(*flights.FlightLeg)

		log.Printf("%s %s->%s\n", flightLeg.FlightNumber, flightLeg.From(), flightLeg.To())
	}
}

func main() {

	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	parser := flags.NewParser(&Options, flags.HelpFlag|flags.PassDoubleDash|flags.PrintErrors)

	_, err := parser.Parse()

	if err != nil {
		os.Exit(1)
	}
}
