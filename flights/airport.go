package flights

import "fmt"

type Airport struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	City      string `json:"city"`
	CityID    string `json:"city_id"`
	Country   string `json:"country"`
	CountryID string `json:"country_id"`
}

func (airport Airport) String() string {

	if airport.Name == "" {

		return fmt.Sprintf("%s  %s, %s", airport.Code, airport.City, airport.Country)
	}

	return fmt.Sprintf("%s  %s, %s, %s", airport.Code, airport.Name, airport.City, airport.Country)
}
