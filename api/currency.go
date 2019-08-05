package api

import (
	"encoding/json"
	"eurogo/shared"
)

// APIBaseURL is the base URL to the exchange rates API.
const APIBaseURL string = "https://api.exchangeratesapi.io/latest"

type CurrencyRates map[string]float64

func GetCurrencyRates(base string) CurrencyRates {

	req, _ := NewRequest(APIBaseURL)
	req.Set("base", base)

	responseChannel := make(chan Response)

	go req.Get(responseChannel)

	response := <-responseChannel

	shared.ErrorHandler(response.Error)

	var currencyResponse struct {
		Rates CurrencyRates `json:"rates"`
		Base  string        `json:"base"`
		Date  string        `json:"date"`
	}

	err := json.Unmarshal([]byte(response.Body), &currencyResponse)
	shared.ErrorHandler(err)

	return currencyResponse.Rates
}

var CurrencySymbols = map[string]string{
	"CAD": "$",
	"HKD": "$",
	"ISK": "kr",
	"PHP": "₱",
	"DKK": "kr",
	"HUF": "Ft",
	"CZK": "Kč",
	"GBP": "£",
	"RON": "lei",
	"SEK": "kr",
	"IDR": "Rp",
	"INR": "₹",
	"BRL": "R$",
	"RUB": "₽",
	"HRK": "kn",
	"JPY": "¥",
	"THB": "฿",
	"CHF": "CHF",
	"EUR": "€",
	"MYR": "RM",
	"BGN": "лв",
	"TRY": "₺",
	"CNY": "¥",
	"NOK": "kr",
	"NZD": "$",
	"ZAR": "R",
	"USD": "$",
	"MXN": "$",
	"SGD": "$",
	"AUD": "$",
	"ILS": "₪",
	"KRW": "₩",
	"PLN": "zł",
}
