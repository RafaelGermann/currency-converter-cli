package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Data map[string]struct {
		Code  string  `json:"code"`
		Value float64 `json:"value"`
	} `json:"data"`
}

func ConvertCurrency(from, to string, amount float64, apiKey string) (float64, error) {
	url := fmt.Sprintf("https://api.currencyapi.com/v3/latest?apikey=%scurrencies=%s&base_currency=%s", apiKey, to, from)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("error fetching data %v", err)
	}
	defer resp.Body.Close()
	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, fmt.Errorf("error decoding response %v", err)
	}

	rate, ok := response.Data[to]
	if !ok {
		return 0, fmt.Errorf("invalid currency: %s", to)
	}

	return amount * rate.Value, nil
}
