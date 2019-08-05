package commands

var Parameters struct {
	// Options
	Currency string `long:"currency" description:"The currency to display prices in."`

	// Commands
	Airport AirportCommand `command:"airport" description:"Search for an airport."`
	Flight  FlightCommand  `command:"flight" description:"Search for an one-way flight over a specified range of dates."`
}
