package services

import (
	"fmt"
)

type Portfolio struct {
	ID      string             `json:"owner-id"`
	Name    string             `json:"name"`
	Stocks  map[string]float32 `json:"stocks"`
	Balance float32            `json:"balance"`
}
type CurrPrice struct {
	Ticker   string  `json:"ticker"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Exchange string  `json:"exchange"`
	Updated  int     `json:"updated"`
	Currency string  `json:"currency"`
}

type HistoricStockPrices struct {
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
	Time   int64   `json:"time"`
}

type User struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Portfolios []Portfolio `json:"portfolios"`
	Email      string      `json:"email"`
	Password   string      `json:"password"`
}

func (c CurrPrice) ToString() {
	fmt.Printf("%+v\n", c)
}
