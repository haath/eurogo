package flights

type AirportsProvider interface {
	SearchAirports(query string) ([]*Airport, error)
}
