package main

import "fmt"

func main() {
	n := -5
	if n > 0 {
		fmt.Println("Positive")
	} else if n < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}

	// Check if a number is even or odd
	num := 7
	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}
}