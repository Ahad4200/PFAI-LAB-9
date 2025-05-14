package main

import (
	"errors"
	"fmt"
)

// average calculates the mean of a slice of float64 values
func average(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("cannot calculate average of empty slice")
	}

	sum := 0.0
	for _, num := range numbers {
		sum += num
	}

	return sum / float64(len(numbers)), nil
}

func main() {
	fmt.Println("Slice Average Calculator:")

	// Test with non-empty slice
	nums := []float64{5.2, 6.8, 9.1}
	avg, err := average(nums)
	if err == nil {
		fmt.Printf("Average of %v: %.2f\n", nums, avg)
	} else {
		fmt.Printf("Error: %s\n", err)
	}

	// Test with empty slice
	emptySlice := []float64{}
	avg, err = average(emptySlice)
	if err == nil {
		fmt.Printf("Average of %v: %.2f\n", emptySlice, avg)
	} else {
		fmt.Printf("Error: %s\n", err)
	}
}
