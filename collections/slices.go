package collections

import (
	"fmt"
)

func Slices() {
	// Declare and initialize a slice
	numbers := []int{10, 20, 30, 40, 50} // Slice with initial values
	fmt.Println("Slice:", numbers)

	// Accessing slice elements
	fmt.Println("First element:", numbers[0])
	fmt.Println("Last element:", numbers[len(numbers)-1])

	// Slicing a slice
	subSlice := numbers[1:4] // Includes elements from index 1 to 3 (excludes 4)
	fmt.Println("Sub-slice [1:4]:", subSlice)

	// Slicing with omitted indices
	fmt.Println("Sub-slice [2:]:", numbers[2:]) // From index 2 to the end
	fmt.Println("Sub-slice [:3]:", numbers[:3]) // From the beginning to index 2

	// Creating a slice using make()
	createdSlice := make([]int, 5) // Creates a slice with length 5, capacity 5
	fmt.Println("Slice created with make():", createdSlice)

	// Append elements to a slice
	fmt.Println("\nAppending elements:")
	createdSlice = append(createdSlice, 60, 70)
	fmt.Println("After appending:", createdSlice)

	// Copying slices
	sourceSlice := []int{1, 2, 3, 4}
	destSlice := make([]int, len(sourceSlice))
	copy(destSlice, sourceSlice) // Copy elements from sourceSlice to destSlice
	fmt.Println("\nSource Slice:", sourceSlice)
	fmt.Println("Destination Slice after copy:", destSlice)

	// Iterating over a slice using range
	fmt.Println("\nIterating over a slice using range:")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Slices are reference types
	fmt.Println("\nSlices are reference types:")
	referenceSlice := numbers
	referenceSlice[0] = 100
	fmt.Println("Original slice after modification:", numbers)
	fmt.Println("Reference slice:", referenceSlice)

	// Capacity and length of a slice
	fmt.Println("\nCapacity and Length:")
	fmt.Printf("Length of numbers slice: %d\n", len(numbers))
	fmt.Printf("Capacity of numbers slice: %d\n", cap(numbers))

	// Expanding slice capacity
	fmt.Println("\nExpanding slice capacity:")
	expandedSlice := append(numbers, 60, 70, 80)
	fmt.Println("Expanded slice:", expandedSlice)
	fmt.Printf("Original slice capacity: %d, Expanded slice capacity: %d\n", cap(numbers), cap(expandedSlice))
}
