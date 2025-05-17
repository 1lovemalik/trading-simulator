package services

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

func (u User) newPortfolio(name string) *Portfolio {
	p := Portfolio{
		ID:      uuid.NewString(),
		Name:    name,
		Stocks:  make(map[string]float32),
		Balance: 0,
	}
	return &p
}

func (user User) AddPortfolio(name string) {
	p := user.newPortfolio(name)

	user.Portfolios = append(user.Portfolios, *p)
}

func CreateUser(newUser User) (*User, error) {
	if newUser.Name == "" || newUser.Email == "" || newUser.Password == "" {
		return nil, fmt.Errorf("failed to create new user. Name: %v, Email: %v, Password: %v",
			newUser.Name, newUser.Email, newUser.Password)
	}

	u := User{
		ID:         uuid.NewString(),
		Name:       newUser.Name,
		Portfolios: make([]Portfolio, 0),
		Email:      newUser.Email,
		Password:   newUser.Password,
	}

	return &u, nil
}

func GetCurrStockData(JsonResponse []byte) (*CurrPrice, error) {

	var currPriceStruct CurrPrice

	err := json.Unmarshal(JsonResponse, &currPriceStruct)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json object. Err: %v", err)
	}

	currPriceStruct.ToString()
	return &currPriceStruct, nil
}
