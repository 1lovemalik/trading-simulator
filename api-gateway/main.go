package main

import (
	"fmt"
	"log"
	"os"

	"trading-simulator/api-gateway/clients"

	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()

	input, err := clients.TakeInputs()
	if err != nil {
		fmt.Println("This is Bad input!")
		clients.TakeInputs()
	}

	_, err = clients.MakeCurrPriceReq(input)
	if err != nil {
		log.Printf("failed to Get Current Price of stock. Err: %v", err)
	}
}

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	clients.NinjaAPIKey = os.Getenv("API_NINJAS_API_KEY")
}
