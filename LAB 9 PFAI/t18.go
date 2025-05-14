package main

import "fmt"

func sumArray(arr []int, ch chan int) {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	ch <- sum
}

func main() {
	// Create two arrays
	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{6, 7, 8, 9, 10}

	// Create channels for results
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Calculate sum concurrently
	go sumArray(array1, ch1)
	go sumArray(array2, ch2)

	// Receive results
	sum1 := <-ch1
	sum2 := <-ch2

	// Print results
	fmt.Printf("Sum of array1 %v: %d\n", array1, sum1)
	fmt.Printf("Sum of array2 %v: %d\n", array2, sum2)
	fmt.Printf("Total sum: %d\n", sum1+sum2)
}
