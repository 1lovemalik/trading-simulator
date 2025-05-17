package main

import (
	"fmt"
	"log"
	"os"
	"trading-simulator/api-gateway/database"
	"trading-simulator/api-gateway/services"

	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()

	err := SetupRepositories()
	if err != nil {
		log.Fatal(err.Error())
	}

	running := true

	for running {

		input, err := services.TakeInputs()
		if err != nil {
			fmt.Println("This is Bad input!")
			services.TakeInputs()
		}

		_, err = services.MakeCurrPriceReq(input)
		if err != nil {
			log.Printf("failed to Get Current Price of stock. Err: %v", err)
		}

	}

}

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	services.NinjaAPIKey = os.Getenv("API_NINJAS_API_KEY")
}

func SetupRepositories() error {

	_, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to db, err: ", err)
	}

	err = database.DBInit()
	if err != nil {
		log.Fatalf("failed to create/load db, err: ", err)
	}

	_, err = database.InitRedis()
	if err != nil {
		log.Fatalf("failed to set up redis client, err: %v", err)
	}

	return nil
}
