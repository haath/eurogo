package flights

type AirportsProvider interface {
	SearchAirports(query string, airports chan<- []*Airport)

	SearchAirportsSync(query string) []*Airport
}
