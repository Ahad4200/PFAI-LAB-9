package main

import (
	"fmt"
	"time"
)

func printNumbers(ch chan int) {
	for i := 1; i <= 3; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func sumArray(arr []int, ch chan int) {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	ch <- sum
}

func main() {
	ch := make(chan int)
	go printNumbers(ch)
	for num := range ch {
		fmt.Println("Received:", num)
	}

	// Create two goroutines to calculate sum of two arrays
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{6, 7, 8, 9, 10}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go sumArray(arr1, ch1)
	go sumArray(arr2, ch2)

	sum1 := <-ch1
	sum2 := <-ch2

	fmt.Printf("Sum of array1: %d\n", sum1)
	fmt.Printf("Sum of array2: %d\n", sum2)
	fmt.Printf("Total sum: %d\n", sum1+sum2)
}
