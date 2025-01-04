package collections

import (
	"fmt"
)

// Define a struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
	City      string
}

// Method with a struct as a receiver
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, %s %s from %s!", p.FirstName, p.LastName, p.City)
}

// Method to update a field using a pointer receiver
func (p *Person) HaveBirthday() {
	p.Age++
}

// Method with additional parameters
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// Define another struct
type Rectangle struct {
	Length, Width float64
}

// Method to calculate the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Method to calculate the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func Collections() {
	// Create an instance of Person
	person1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
		City:      "New York",
	}
	fmt.Println("Person Details:", person1)

	// Call the Greet method
	fmt.Println("\nGreeting:", person1.Greet())

	// Call the HaveBirthday method
	person1.HaveBirthday()
	fmt.Printf("After birthday, %s's age is: %d\n", person1.FirstName, person1.Age)

	// Call the IsAdult method
	if person1.IsAdult() {
		fmt.Println(person1.FirstName, "is an adult.")
	} else {
		fmt.Println(person1.FirstName, "is not an adult.")
	}

	// Create an instance of Rectangle
	rect := Rectangle{
		Length: 5.0,
		Width:  3.0,
	}
	fmt.Println("\nRectangle Details:", rect)

	// Call methods on Rectangle
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

	// Anonymous struct
	anonymousPerson := struct {
		Name string
		Job  string
	}{
		Name: "Alice",
		Job:  "Engineer",
	}
	fmt.Println("\nAnonymous Struct:", anonymousPerson)
}
