package main

import "fmt"

func main() {
	var x int = 5
	const y = 10
	x += 1

	var name string = "GoLang Programmer" // Added string variable
	fmt.Printf("x = %d, y = %d\n", x, y)
	fmt.Printf("Name: %s\n", name) // Printing the string
}
