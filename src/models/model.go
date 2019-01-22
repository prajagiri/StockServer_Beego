package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	url       = "https://www.worldtradingdata.com/api/v1/stock?symbol="
	api_token = "5LM183iQ3mKUWX8s6txwwfGX1Qtz0WU7ekLo6zLf1LJM9gIx9MWJ6PI5S06Z"
)

type StockData struct {
	Symbols_requested int                      `json:"symbols_requested"`
	Symbols_returned  int                      `json:"symbols_returned"`
	Data              []map[string]interface{} `json:"data"`
}

func GetStockData(symbol, stock_exchange string) (StockData, error) {
	var formURL string
	if len(stock_exchange) > 0 {
		formURL = url + symbol + "&stock_exchange=" + stock_exchange + "&api_token=" + api_token
	} else {
		formURL = url + symbol + "&stock_exchange=AMEX" + "&api_token=" + api_token
	}

	fmt.Println("Test - ", formURL)
	resp, err := http.Get(formURL)
	if err != nil {
		log.Fatalln(err)
		return StockData{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return StockData{}, err
	}

	var data StockData
	err = json.Unmarshal(body, &data)

	return data, nil
}
