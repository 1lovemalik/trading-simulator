package main

import (
	"fmt"
	"log"
)

func main() {
	loadEnv()

	input, err := takeInputs()
	if err != nil {
		fmt.Println("This is Bad input!")
		takeInputs()
	}

	_, err = MakeCurrPriceReq(input)
	if err != nil {
		log.Printf("failed to Get Current Price of stock. Err: %v", err)
	}
}
