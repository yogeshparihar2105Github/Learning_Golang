package error_handling

import (
	"fmt"
)

// Custom error type for validation errors
type CustomValidationError struct {
	Field   string
	Message string
}

// Implement the Error() method for the ValidationError type
func (e *CustomValidationError) Error() string {
	return fmt.Sprintf("Validation error on field '%s': %s", e.Field, e.Message)
}

// Custom error type for database errors
type DatabaseError struct {
	Code    int
	Message string
}

// Implement the Error() method for the DatabaseError type
func (e *DatabaseError) Error() string {
	return fmt.Sprintf("Database error (Code %d): %s", e.Code, e.Message)
}

// Function that returns a validation error
func customValidateInput(age int) error {
	if age < 0 {
		return &CustomValidationError{
			Field:   "age",
			Message: "must be a non-negative number",
		}
	}
	return nil
}

// Function that returns a database error
func fetchFromDatabase(id int) error {
	if id <= 0 {
		return &DatabaseError{
			Code:    404,
			Message: "record not found",
		}
	}
	return nil
}

func CustomErrors() {
	fmt.Println("Custom Error Handling Examples")

	// Example: Handling a validation error
	fmt.Println("\nHandling ValidationError:")
	err := customValidateInput(-5)
	if err != nil {
		if ve, ok := err.(*CustomValidationError); ok {
			fmt.Printf("Validation Error: Field=%s, Message=%s\n", ve.Field, ve.Message)
		} else {
			fmt.Println("Unexpected error:", err)
		}
	} else {
		fmt.Println("Input is valid.")
	}

	// Example: Handling a database error
	fmt.Println("\nHandling DatabaseError:")
	err = fetchFromDatabase(0)
	if err != nil {
		if de, ok := err.(*DatabaseError); ok {
			fmt.Printf("Database Error: Code=%d, Message=%s\n", de.Code, de.Message)
		} else {
			fmt.Println("Unexpected error:", err)
		}
	} else {
		fmt.Println("Database fetch succeeded.")
	}

	// Example: No error case
	fmt.Println("\nHandling No Error Case:")
	err = customValidateInput(25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Input is valid.")
	}
}
