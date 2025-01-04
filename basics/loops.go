package basics

import (
	"fmt"
)

func Loops() {
	// Basic For Loop
	fmt.Println("Basic For Loop:")
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// For Loop with Only Condition (like a while loop)
	fmt.Println("\nFor Loop with Only Condition:")
	counter := 0
	for counter < 3 {
		fmt.Println("Counter:", counter)
		counter++
	}

	// Infinite Loop with Break
	fmt.Println("\nInfinite Loop with Break:")
	sum := 0
	for {
		sum++
		if sum == 5 {
			fmt.Println("Sum reached:", sum)
			break
		}
	}

	// For Loop with Continue
	fmt.Println("\nFor Loop with Continue:")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println("Odd number:", i)
	}

	// Iterating Over an Array with Range
	fmt.Println("\nIterating Over an Array:")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Iterating Over a Map with Range
	fmt.Println("\nIterating Over a Map:")
	person := map[string]string{
		"name":    "John",
		"age":     "30",
		"country": "USA",
	}
	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Iterating Over a String (Unicode Characters)
	fmt.Println("\nIterating Over a String:")
	text := "Hello, 世界"
	for index, char := range text {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	// For Loop Without Range
	fmt.Println("\nFor Loop Without Range:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index: %d, Value: %d\n", i, numbers[i])
	}
}
