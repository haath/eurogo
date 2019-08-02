package skyscanner

import "eurogo/flights"

type skyscannerAirportsResponse []skyscannerAirport

type skyscannerAirport struct {
	PlaceID   string `json:"PlaceId"`
	City      string `json:"CityName"`
	CityID    string `json:"CityId"`
	Country   string `json:"CountryName"`
	CountryID string `json:"CountryId"`
}

func (resp *skyscannerAirportsResponse) getAirports() []flights.Airport {

	var airports []flights.Airport

	for _, airportData := range *resp {

		airport := flights.Airport{
			Code:      airportData.PlaceID,
			City:      airportData.City,
			CityID:    airportData.CityID,
			Country:   airportData.Country,
			CountryID: airportData.CountryID,
		}

		airports = append(airports, airport)
	}

	return airports
}
