package main

import "fmt"

func main() {
	a := 10
	b := 3
	fmt.Printf("Sum: %d\n", a+b)
	fmt.Printf("Product: %d\n", a*b)

	// Calculate remainder of 15 divided by 4
	remainder := 15 % 4
	fmt.Printf("Remainder of 15/4: %d\n", remainder)
}