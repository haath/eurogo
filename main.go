package main

import (
	"log"
	"os"

	"eurogo/commands"

	"github.com/jessevdk/go-flags"
)

func main() {

	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	commands.Parameters.Currency = "EUR"

	parser := flags.NewParser(&commands.Parameters, flags.HelpFlag|flags.PassDoubleDash|flags.PrintErrors)

	_, err := parser.Parse()

	if err != nil {
		os.Exit(1)
	}
}
