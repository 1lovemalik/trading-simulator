package main

import "fmt"

var CommandMap = map[string]string{
	"Current Price": "https://api.api-ninjas.com/v1/stockprice?ticker=",
}

type CurrPrice struct {
	Ticker   string  `json:"ticker"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Exchange string  `json:"exchange"`
	Updated  int     `json:"updated"`
	Currency string  `json:"currency"`
}

func (c CurrPrice) ToString() {
	fmt.Printf("%+v\n", c)
}
