package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	_ "github.com/joho/godotenv"
)

var Client = &http.Client{Timeout: 10 * time.Second}
var NinjaAPIKey string

func takeInputs() (string, error) {
	fmt.Print("Enter the ticker symbol of the stock that you want!")

	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return "", fmt.Errorf("bad Input! Error: %v", err)
	}

	InputChecker(input)

	url, exists := CommandMap["Current Price"]
	if !exists {
		url = "https://api.api-ninjas.com/v1/stockprice?ticker="
		CommandMap["Current Price"] = url
	}
	url += input
	return url, nil
}

func MakeCurrPriceReq(url string) (*CurrPrice, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}

	req.Header.Add("X-API-KEY", NinjaAPIKey)
	response, err := Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to complete request. Err: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-OK status: %d, %s", response.StatusCode, response.Status)
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var currPriceStruct CurrPrice

	err = json.Unmarshal(resBody, &currPriceStruct)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json object. Err: %v", err)
	}

	currPriceStruct.ToString()
	return &currPriceStruct, nil
}

func getHistoricalStockPrices(tickerName string) ([]HistoricStockPrices, error) {
	err := InputChecker(tickerName)
	if err != nil {
		return nil, fmt.Errorf("Bad Ticker Name. Err: %s", err)
	}

	url := fmt.Sprintf(
		"https://api.api-ninjas.com/v1/stockpricehistorical?ticker=%s&period=1h",
		tickerName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to make new request. Err: %s", err)
	}

	req.Header.Add("X-API-KEY", NinjaAPIKey)

	response, err := Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute get request. Err: %s", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bad status code; %v, Body: %s", response.StatusCode, response.Body)
	}

	jsonArray, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from response Body: %s", err)
	}

	var PricesArray []HistoricStockPrices

	err = json.Unmarshal(jsonArray, &PricesArray)
	if err != nil {
		return nil, fmt.Errorf("Failed to Unmarshal from Json. Err: %s", err)
	}

	return PricesArray, nil
}
