package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	_ "github.com/joho/godotenv"
)

var client = &http.Client{Timeout: 10 * time.Second}
var NinjaAPIKey string

func takeInputs() (string, error) {
	fmt.Print("Enter the ticker symbol of the stock that you want!")

	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return "", fmt.Errorf("bad Input! Error: %v", err)
	}

	url, exists := CommandMap["Current Price"]
	if !exists {
		url = "https://api.api-ninjas.com/v1/stockprice?ticker="
		CommandMap["Current Price"] = url
	}
	url += input
	return url, nil
}

func MakeCurrPriceReq(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to make request: %v", err)
	}

	req.Header.Add("X-API-KEY", NinjaAPIKey)
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to complete request. Err: %v", err)

	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	var currPriceStruct CurrPrice

	err = json.Unmarshal(resBody, &currPriceStruct)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json object. Err: %v", err)
	}

	currPriceStruct.ToString()
	return nil
}
