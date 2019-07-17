package main

import (
	"log"
	"os"
	"time"

	"eurogo/flights/skiplagged"

	"github.com/jessevdk/go-flags"
)

var Options struct {
	Call func(n int) `short:"c"`
}

func foo() {

	prov := skiplagged.SkiplaggedFlightProvider()

	log.Println(prov.Search("PRG", "ATH", time.Date(2019, time.October, 12, 0, 0, 0, 0, time.UTC)))
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
}
