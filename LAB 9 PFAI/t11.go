package main

import (
	"fmt"
	"os"
)

func main() {
	os.WriteFile("data.txt", []byte("Go is efficient!"), 0644)
	data, _ := os.ReadFile("data.txt")
	fmt.Println("File content:", string(data))

	// Append a new line to the file
	file, _ := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString("\nGo is also simple and powerful!")

	// Read updated content
	updatedData, _ := os.ReadFile("data.txt")
	fmt.Println("Updated content:", string(updatedData))
}
