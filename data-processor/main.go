package main

import "fmt"

func main() {
	loadEnv()

	input, err := takeInputs()
	if err != nil {
		fmt.Println("This is Bad input!")
		takeInputs()
	}

	MakeCurrPriceReq(input)
}
