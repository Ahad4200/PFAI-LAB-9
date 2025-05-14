package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	dict := map[string]string{"name": "Alice", "job": "Engineer"}
	fmt.Println("Job:", dict["job"])

	// Add a new key-value pair
	dict["location"] = "San Francisco"
	fmt.Println("Location:", dict["location"])

	p := Person{Name: "Alice", Age: 30}
	fmt.Printf("Person: %s, Age: %d\n", p.Name, p.Age)
}
