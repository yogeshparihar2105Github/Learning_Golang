package basics

import (
	"fmt"
)

func Conditionals() {
	// Example: If-Else Statement
	num := 42

	fmt.Println("If-Else Example:")
	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// If with an initialization statement
	fmt.Println("\nIf with Initialization:")
	if value := 100; value > 50 {
		fmt.Println("The value", value, "is greater than 50.")
	} else {
		fmt.Println("The value", value, "is not greater than 50.")
	}

	// Nested If-Else
	fmt.Println("\nNested If-Else:")
	age := 20
	if age >= 18 {
		if age >= 21 {
			fmt.Println("You are an adult and can legally drink in the US.")
		} else {
			fmt.Println("You are an adult but cannot legally drink in the US.")
		}
	} else {
		fmt.Println("You are not an adult.")
	}

	// Example: Switch Statement
	day := 3

	fmt.Println("\nSwitch Example:")
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend")
	}

	// Switch with multiple cases
	fmt.Println("\nSwitch with Multiple Cases:")
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("It's a weekday.")
	case 6, 7:
		fmt.Println("It's the weekend.")
	default:
		fmt.Println("Invalid day.")
	}

	// Switch without a condition (acts like an if-else chain)
	number := -5

	fmt.Println("\nSwitch Without Condition:")
	switch {
	case number > 0:
		fmt.Println(number, "is positive.")
	case number < 0:
		fmt.Println(number, "is negative.")
	default:
		fmt.Println(number, "is zero.")
	}

	// Switch with fallthrough (forces execution of the next case)
	fmt.Println("\nSwitch with Fallthrough:")
	grade := "B"
	switch grade {
	case "A":
		fmt.Println("Excellent!")
	case "B":
		fmt.Println("Good job!")
		fallthrough
	case "C":
		fmt.Println("You passed.")
	default:
		fmt.Println("Better luck next time.")
	}
}
