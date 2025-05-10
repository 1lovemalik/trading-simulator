package main

import (
	"fmt"
	"unicode"
)

func InputChecker(input string) error {
	if len(input) == 0 {
		return fmt.Errorf("Input cannot be empty!")
	}

	if containsNumbers(input) {
		return fmt.Errorf("A ticker symbol cannot contain numbers!")
	}

	return nil
}

func containsNumbers(input string) bool {
	chars := []byte(input)

	for i := 0; i < len(chars); i++ {
		if unicode.IsDigit(rune(chars[i])) {
			return true
		}
	}
	return false
}
