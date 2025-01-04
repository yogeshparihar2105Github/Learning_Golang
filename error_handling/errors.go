package error_handling

import (
	"errors"
	"fmt"
)

// Function that returns a basic error
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Custom error using fmt.Errorf
func checkEven(number int) error {
	if number%2 != 0 {
		return fmt.Errorf("the number %d is not even", number)
	}
	return nil
}

// Custom error type
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s - %s", e.Field, e.Message)
}

// Function that uses a custom error type
func validateInput(age int) error {
	if age < 0 {
		return &ValidationError{
			Field:   "age",
			Message: "must be a non-negative value",
		}
	}
	return nil
}

func ErrorHandling() {
	fmt.Println("Error Handling Examples")

	// Example: Basic error handling
	fmt.Println("\nBasic Error Handling:")
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Example: Using fmt.Errorf
	fmt.Println("\nUsing fmt.Errorf:")
	if err := checkEven(5); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The number is even.")
	}

	// Example: Custom error type
	fmt.Println("\nCustom Error Type:")
	if err := validateInput(-1); err != nil {
		switch e := err.(type) {
		case *ValidationError:
			fmt.Printf("Custom error caught: Field=%s, Message=%s\n", e.Field, e.Message)
		default:
			fmt.Println("Unexpected error:", err)
		}
	} else {
		fmt.Println("Input is valid.")
	}

	// Example: Handling a valid case
	fmt.Println("\nValid Case:")
	result, err = divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
