package routes

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Portfolios []Portfolio `json:"portfolios"`
	Email      string      `json:"email"`
	Password   string      `json:"password"`
}

func (u User) addPortfolio(name string) {
	p := u.newPortfolio(name)

	u.Portfolios = append(u.Portfolios, *p)
}

func createUser(newUser User) (*User, error) {
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
