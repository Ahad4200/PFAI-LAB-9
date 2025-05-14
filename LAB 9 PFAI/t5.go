package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	counter := 1
	for counter <= 3 {
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}

	// Print numbers from 10 to 1 in reverse
	fmt.Println("Counting backward:")
	for i := 10; i >= 1; i-- {
		fmt.Printf("%d ", i)
	}
	fmt.Println() // New line after the countdown
}