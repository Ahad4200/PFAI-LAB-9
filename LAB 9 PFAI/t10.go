package main

import (
	"fmt"
	"strconv"
)

func parseNumber(s string) (int, error) {
	return strconv.Atoi(s)
}

func main() {
	if num, err := parseNumber("123"); err == nil {
		fmt.Println("Number:", num)
	} else {
		fmt.Println("Error:", err)
	}

	// Handle invalid input
	if num, err := parseNumber("abc"); err == nil {
		fmt.Println("Number:", num)
	} else {
		fmt.Println("Error:", err)
	}
}
