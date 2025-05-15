package routes

import "github.com/google/uuid"

type Portfolio struct {
	ID      string             `json:"owner-id"`
	Name    string             `json:"name"`
	Stocks  map[string]float32 `json:"stocks"`
	Balance float32            `json:"balance"`
}

func (u User) newPortfolio(name string) *Portfolio {
	p := Portfolio{
		ID:      uuid.NewString(),
		Name:    name,
		Stocks:  make(map[string]float32),
		Balance: 0,
	}
	return &p
}
