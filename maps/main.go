package main

import "fmt"

func main() {
	// Declare and initialize a map
	studentMarks := map[string]int{
		"Alice": 90,
		"Bob":   85,
		"Eve":   92,
	}

	// Access a value
	fmt.Println("Alice's marks:", studentMarks["Alice"])

	// Add a new key-value pair
	studentMarks["Charlie"] = 88

	// Update a value
	studentMarks["Bob"] = 89

	// Delete a key
	delete(studentMarks, "Eve")

	// Check if a key exists
	mark, exists := studentMarks["Eve"]
	if exists {
		fmt.Println("Eve's marks:", mark)
	} else {
		fmt.Println("Eve not found")
	}

	// Iterate over map
	for name, mark := range studentMarks {
		fmt.Printf("%s got %d\n", name, mark)
	}
}
