package collections

import (
	"fmt"
)

func Maps() {
	// Declare and initialize a map
	person := map[string]string{
		"name":    "John",
		"age":     "30",
		"country": "USA",
	}
	fmt.Println("Initial map:", person)

	// Accessing map values
	fmt.Println("Name:", person["name"])
	fmt.Println("Age:", person["age"])

	// Adding and updating map values
	person["city"] = "New York" // Adding a new key-value pair
	person["age"] = "31"        // Updating an existing key-value pair
	fmt.Println("\nUpdated map:", person)

	// Deleting a key-value pair
	delete(person, "country")
	fmt.Println("\nMap after deleting 'country':", person)

	// Checking if a key exists
	value, exists := person["country"]
	if exists {
		fmt.Println("\n'country' exists with value:", value)
	} else {
		fmt.Println("\n'country' does not exist.")
	}

	// Iterating over a map using range
	fmt.Println("\nIterating over a map:")
	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Creating a map using make()
	cities := make(map[string]int)
	cities["New York"] = 8419600
	cities["Los Angeles"] = 3980400
	cities["Chicago"] = 2716000
	fmt.Println("\nCities map:", cities)

	// Accessing a non-existent key
	population := cities["San Francisco"]
	fmt.Println("\nPopulation of San Francisco (non-existent key):", population)

	// Length of a map
	fmt.Println("\nLength of the cities map:", len(cities))

	// Nested maps
	nestedMap := map[string]map[string]string{
		"person1": {
			"name": "Alice",
			"age":  "25",
		},
		"person2": {
			"name": "Bob",
			"age":  "30",
		},
	}
	fmt.Println("\nNested map:", nestedMap)

	// Accessing nested map values
	fmt.Println("Name of person2:", nestedMap["person2"]["name"])
}
