package collections

import (
	"fmt"
)

func Arrays() {
	// Declare and initialize an array
	var numbers [5]int                            // Array of integers with length 5
	fmt.Println("Default array values:", numbers) // Default values are zero

	// Assign values to array elements
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println("Updated array:", numbers)

	// Declare and initialize an array in a single line
	colors := [3]string{"Red", "Green", "Blue"}
	fmt.Println("Colors array:", colors)

	// Using ellipsis to let Go determine the size of the array
	prices := [...]float64{10.5, 20.0, 30.25}
	fmt.Println("Prices array:", prices)

	// Accessing array elements
	fmt.Println("First element of numbers array:", numbers[0])
	fmt.Println("Last element of colors array:", colors[len(colors)-1])

	// Iterating over an array using a for loop
	fmt.Println("\nIterating using a for loop:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
	}

	// Iterating over an array using range
	fmt.Println("\nIterating using range:")
	for index, value := range colors {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// Multidimensional arrays
	var matrix [2][2]int // 2x2 array
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[1][0] = 3
	matrix[1][1] = 4
	fmt.Println("\n2x2 Matrix:")
	fmt.Println(matrix)

	// Accessing elements of a multidimensional array
	fmt.Println("Element at [1][1]:", matrix[1][1])

	// Arrays are value types in Go (copy behavior)
	copyArray := numbers
	copyArray[0] = 100
	fmt.Println("\nOriginal array:", numbers)
	fmt.Println("Copied array after modification:", copyArray)
}
