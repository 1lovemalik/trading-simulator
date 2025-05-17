package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"trading-simulator/api-gateway/util"

	_ "github.com/joho/godotenv"
)

var CommandMap = map[string]string{
	"Current Price": "https://api.api-ninjas.com/v1/stockprice?ticker=",
}

func TakeInputs() (string, error) {
	fmt.Println("Enter the ticker symbol of the stock that you want!")

	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return "", fmt.Errorf("bad Input! Error: %v", err)
	}

	err = util.InputChecker(input)
	if err != nil {
		return "", fmt.Errorf("this is not valid input!: err: %v", err)
	}

	return input, nil
}

var Client = &http.Client{Timeout: 10 * time.Second}
var NinjaAPIKey string

func MakeCurrPriceReq(input string) (*CurrPrice, error) {
	url := fmt.Sprintf("https://api.api-ninjas.com/v1/stockprice?ticker=%v", input)
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

	prices, err := GetCurrStockData(resBody)
	if err != nil {
		return nil, fmt.Errorf("get current stock data service raised err: %v", err)
	}
	return prices, nil
}

func GetHistoricalStockPrices(tickerName string) ([]HistoricStockPrices, error) {
	err := util.InputChecker(tickerName)
	if err != nil {
		return nil, fmt.Errorf("Bad Ticker Name. Err: %s", err)
	}

	url := fmt.Sprintf(
		"https://api.api-ninjas.com/v1/stockpricehistorical?ticker=%s&period=1h",
		tickerName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make new request. Err: %s", err)
	}

	req.Header.Add("X-API-KEY", NinjaAPIKey)

	response, err := Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute get request. Err: %s", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code; %v, Body: %s", response.StatusCode, response.Body)
	}

	jsonArray, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read from response Body: %s", err)
	}

	var PricesArray []HistoricStockPrices

	err = json.Unmarshal(jsonArray, &PricesArray)
	if err != nil {
		return nil, fmt.Errorf("failed to Unmarshal from Json. Err: %s", err)
	}

	return PricesArray, nil
}
