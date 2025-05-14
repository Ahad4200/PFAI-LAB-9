package main

import "fmt"

func square(x int) int {
	return x * x
}

// Calculate factorial
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Printf("Square of 5: %d\n", square(5))
	fmt.Printf("Factorial of 5: %d\n", factorial(5))
}