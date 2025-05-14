package main

import "fmt"

func main() {
	arr := [3]int{10, 20, 30}
	slice := append(arr[:], 40)
	fmt.Printf("First element: %d\n", arr[0])
	fmt.Println("Slice:", slice)

	// Create and iterate over a slice of strings
	fruits := []string{"Apple", "Orange", "Banana", "Mango"}
	fmt.Println("Fruits:")
	for i, fruit := range fruits {
		fmt.Printf("%d: %s\n", i, fruit)
	}
}
