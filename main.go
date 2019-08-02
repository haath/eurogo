package main

import (
	"log"
	"os"
	"time"

	"eurogo/flights"
	"eurogo/flights/skiplagged"

	"github.com/jessevdk/go-flags"
)

var Options struct {
	Call func(n int) `short:"c"`
}

func foo() {

	prov := skiplagged.SkiplaggedFlightProvider()

	flightsList, err := prov.Search("PRG", "ATH", time.Date(2019, time.October, 12, 0, 0, 0, 0, time.UTC))

	if err != nil {
		log.Fatal(err)
	}

	flight := flightsList[0]
	for _, leg := range flight.Legs {

		flightLeg := leg.(*flights.FlightLeg)
	}
}

func main() {

	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	Options.Call = func(n int) {
		log.Println(n)
	}

	parser := flags.NewParser(&Options, flags.HelpFlag|flags.PassDoubleDash|flags.PrintErrors)

	_, err := parser.Parse()

	if err != nil {
		os.Exit(1)
	}

	foo()
}
