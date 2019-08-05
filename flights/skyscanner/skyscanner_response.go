package skyscanner

import (
	"eurogo/flights"
	"eurogo/flights/skiplagged"
)

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
	var airportNameChannels []chan string

	for _, airportData := range *resp {

		airport := flights.Airport{
			Code:      airportData.PlaceID,
			City:      airportData.City,
			CityID:    airportData.CityID,
			Country:   airportData.Country,
			CountryID: airportData.CountryID,
		}

		airportNameChannel := make(chan string)
		airportNameChannels = append(airportNameChannels, airportNameChannel)

		go skiplagged.GetAirportName(airport.Code, airportNameChannel)

		airports = append(airports, airport)
	}

	for i := range airports {

		airports[i].Name = <-airportNameChannels[i]
	}

	return airports
}
