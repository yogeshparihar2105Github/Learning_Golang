package basics

import (
	"fmt"
)

func Variables_Example() {
	// Declare variables using `var`
	var message string = "Hello, World!"
	var number int = 42
	var pi float64 = 3.14159
	var isGoFun bool = true

	// Declare multiple variables in one line
	var a, b, c int = 1, 2, 3

	// Type inference (Go determines the type)
	var inferredString = "Go is awesome!"
	var inferredNumber = 100

	// Zero values (default values for uninitialized variables)
	var uninitializedInt int       // 0
	var uninitializedString string // ""
	var uninitializedBool bool     // false

	// Short variable declaration (only allowed inside functions)
	shortVar := "Short Declaration Example"
	x, y := 10, 20 // Multiple variables

	// Constants
	const constantMessage string = "This is a constant"
	const constantNumber = 3.14 // Type inference works for constants too
	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)

	// Print variables and constants
	fmt.Println("Variables:")
	fmt.Println("message:", message)
	fmt.Println("number:", number)
	fmt.Println("pi:", pi)
	fmt.Println("isGoFun:", isGoFun)
	fmt.Println("a, b, c:", a, b, c)
	fmt.Println("inferredString:", inferredString)
	fmt.Println("inferredNumber:", inferredNumber)
	fmt.Println("uninitializedInt:", uninitializedInt)
	fmt.Println("uninitializedString:", uninitializedString)
	fmt.Println("uninitializedBool:", uninitializedBool)
	fmt.Println("shortVar:", shortVar)
	fmt.Println("x, y:", x, y)

	fmt.Println("\nConstants:")
	fmt.Println("constantMessage:", constantMessage)
	fmt.Println("constantNumber:", constantNumber)
	fmt.Println("Days:", Monday, Tuesday, Wednesday)
}
