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

	airports := provider.SearchAirportsSync(query)

	if Parameters.JSON {

		RenderAirportsJSON(airports)

	} else {

		for _, airport := range airports {

			log.Println(airport)
		}
	}

	return nil
}
