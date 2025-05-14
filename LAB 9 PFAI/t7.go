package main

import "fmt"

func swap(a, b string) (string, string) {
	return b, a
}

// Return both sum and product
func sumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	x, y := swap("hello", "world")
	fmt.Println(x, y)

	sum, product := sumAndProduct(5, 3)
	fmt.Printf("Sum: %d, Product: %d\n", sum, product)
}
