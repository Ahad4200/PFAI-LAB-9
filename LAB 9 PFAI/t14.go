package main

import (
	"fmt"
	"math"
)

// isPrime checks if a number is prime
func isPrime(n int) bool {
	// 0 and 1 are not prime numbers
	if n <= 1 {
		return false
	}

	// 2 is prime
	if n == 2 {
		return true
	}

	// Even numbers (except 2) are not prime
	if n%2 == 0 {
		return false
	}

	// Check odd divisors up to square root of n
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	testCases := []int{0, 1, 2, 13, 27}

	fmt.Println("Prime Number Checker:")
	for _, num := range testCases {
		if isPrime(num) {
			fmt.Printf("%d is prime\n", num)
		} else {
			fmt.Printf("%d is not prime\n", num)
		}
	}
}
