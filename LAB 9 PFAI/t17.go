package main

import (
	"fmt"
	"os"
)

func appendToFile(filename, content string) error {
	// Check if file exists
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// Create file if it doesn't exist
		return os.WriteFile(filename, []byte(content), 0644)
	}

	// Open file for appending
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Append content
	_, err = file.WriteString(content)
	return err
}

func main() {
	filename := "data.txt"
	content := "\nThis is new content appended to the file."

	err := appendToFile(filename, content)
	if err != nil {
		fmt.Printf("Failed to append to file: %s\n", err)
	} else {
		fmt.Printf("Successfully appended content to %s\n", filename)

		// Read and print the file content
		data, err := os.ReadFile(filename)
		if err == nil {
			fmt.Printf("Current file content:\n%s\n", string(data))
		}
	}
}
