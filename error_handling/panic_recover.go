package error_handling

import (
	"fmt"
)

func safeFunction() {
	fmt.Println("Executing safe function.")
}

func riskyFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Executing risky function.")
	panic("Something went wrong!")                 // Trigger a panic
	fmt.Println("This line will not be executed.") // Unreachable
}

func PanicRecover() {
	fmt.Println("PanicRecover function starts.")

	// Call a safe function
	safeFunction()

	// Call a risky function that causes a panic
	fmt.Println("\nCalling risky function:")
	riskyFunction()

	// Program continues after recovering from panic
	fmt.Println("\nProgram continues after recovering from panic.")

	// Additional panics without recover will crash the program
	fmt.Println("\nAnother risky function:")
	func() {
		fmt.Println("Executing another risky function.")
		panic("Fatal error!")
	}()

	fmt.Println("This line will not be executed due to unhandled panic.")
}
