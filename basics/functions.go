package basics

import (
	"fmt"
)

// A simple function with no parameters and no return value
func greet() {
	fmt.Println("Hello, World!")
}

// A function with parameters
func add(a int, b int) int {
	return a + b
}

// A function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

// A function with named return values
func rectangleDimensions(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}

// A function with a variadic parameter
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// A function as a parameter
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Anonymous function assigned to a variable
var multiply = func(a, b int) int {
	return a * b
}

// Recursive function
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func Functions() {
	// Call a simple function
	fmt.Println("Calling greet:")
	greet()

	// Call a function with parameters
	fmt.Println("\nCalling add:")
	fmt.Println("5 + 3 =", add(5, 3))

	// Call a function with multiple return values
	fmt.Println("\nCalling divide:")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	// Call a function with named return values
	fmt.Println("\nCalling rectangleDimensions:")
	area, perimeter := rectangleDimensions(5, 3)
	fmt.Println("Area:", area, "Perimeter:", perimeter)

	// Call a variadic function
	fmt.Println("\nCalling sumAll:")
	fmt.Println("Sum of 1, 2, 3, 4, 5:", sumAll(1, 2, 3, 4, 5))

	// Call a function as a parameter
	fmt.Println("\nCalling applyOperation with multiply:")
	fmt.Println("3 * 4 =", applyOperation(3, 4, multiply))

	// Recursive function
	fmt.Println("\nCalling factorial:")
	fmt.Println("Factorial of 5:", factorial(5))
}
