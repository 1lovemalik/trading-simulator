package routes

type Portfolio struct {
	OwnerID string             `json:"owner-id"`
	Name    string             `json:"name"`
	Stocks  map[string]float32 `json:"stocks"`
	Balance float32            `json"balance":`
}

func (u User) newPortfolio(name string) *Portfolio {
	p := Portfolio{
		OwnerID: u.ID,
		Name:    name,
		Stocks:  make(map[string]float32),
		Balance: 0,
	}
	return &p

}
