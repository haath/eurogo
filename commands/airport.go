package commands

import (
	"eurogo/flights/skyscanner"
	"log"
)

type AirportCommand struct {
	Args airportPositionalArgs `positional-args:"1" required:"1"`
}

type airportPositionalArgs struct {
	Query string `positional-arg-name:"<query>"`
}

func (cmd *AirportCommand) Execute(args []string) error {

	provider := skyscanner.SkyscannerAirportsProvider()

	query := cmd.Args.Query

	airports, err := provider.SearchAirports(query)

	if err != nil {
		return err
	}

	for _, airport := range airports {

		log.Println(airport)
	}

	return nil
}
