package basics

import (
	"fmt"
	"reflect"
)

func DataTypes() {
	// Integer data types
	var intValue int = 42
	var int8Value int8 = 127
	var int16Value int16 = 32767
	var int32Value int32 = 2147483647
	var int64Value int64 = 9223372036854775807

	// Unsigned integer data types
	var uintValue uint = 42
	var uint8Value uint8 = 255
	var uint16Value uint16 = 65535
	var uint32Value uint32 = 4294967295
	var uint64Value uint64 = 18446744073709551615

	// Floating-point data types
	var float32Value float32 = 3.14
	var float64Value float64 = 3.141592653589793

	// Complex numbers
	var complex64Value complex64 = 1 + 2i
	var complex128Value complex128 = 2 + 3i

	// Boolean data type
	var boolValue bool = true

	// String data type
	var stringValue string = "Hello, Go!"

	// Rune (alias for int32, represents a Unicode code point)
	var runeValue rune = 'A' // Unicode for 'A' is 65

	// Byte (alias for uint8)
	var byteValue byte = 'B' // ASCII for 'B' is 66

	// Type inference with shorthand declaration
	autoInt := 50
	autoFloat := 5.67
	autoBool := false
	autoString := "Type Inferred"

	// Type conversion
	var intToFloat float64 = float64(intValue)
	var floatToInt int = int(float32Value)
	var intToString string = fmt.Sprintf("%d", intValue)

	// Print data types and values
	fmt.Println("Integer Types:")
	fmt.Printf("int: %d, Type: %s\n", intValue, reflect.TypeOf(intValue))
	fmt.Printf("int8: %d, Type: %s\n", int8Value, reflect.TypeOf(int8Value))
	fmt.Printf("int16: %d, Type: %s\n", int16Value, reflect.TypeOf(int16Value))
	fmt.Printf("int32: %d, Type: %s\n", int32Value, reflect.TypeOf(int32Value))
	fmt.Printf("int64: %d, Type: %s\n", int64Value, reflect.TypeOf(int64Value))

	fmt.Println("\nUnsigned Integer Types:")
	fmt.Printf("uint: %d, Type: %s\n", uintValue, reflect.TypeOf(uintValue))
	fmt.Printf("uint8: %d, Type: %s\n", uint8Value, reflect.TypeOf(uint8Value))
	fmt.Printf("uint16: %d, Type: %s\n", uint16Value, reflect.TypeOf(uint16Value))
	fmt.Printf("uint32: %d, Type: %s\n", uint32Value, reflect.TypeOf(uint32Value))
	fmt.Printf("uint64: %d, Type: %s\n", uint64Value, reflect.TypeOf(uint64Value))

	fmt.Println("\nFloating-Point Types:")
	fmt.Printf("float32: %f, Type: %s\n", float32Value, reflect.TypeOf(float32Value))
	fmt.Printf("float64: %f, Type: %s\n", float64Value, reflect.TypeOf(float64Value))

	fmt.Println("\nComplex Types:")
	fmt.Printf("complex64: %v, Type: %s\n", complex64Value, reflect.TypeOf(complex64Value))
	fmt.Printf("complex128: %v, Type: %s\n", complex128Value, reflect.TypeOf(complex128Value))

	fmt.Println("\nBoolean Type:")
	fmt.Printf("bool: %t, Type: %s\n", boolValue, reflect.TypeOf(boolValue))

	fmt.Println("\nString Type:")
	fmt.Printf("string: %s, Type: %s\n", stringValue, reflect.TypeOf(stringValue))

	fmt.Println("\nRune and Byte Types:")
	fmt.Printf("rune: %c, Unicode: %U, Type: %s\n", runeValue, runeValue, reflect.TypeOf(runeValue))
	fmt.Printf("byte: %c, ASCII: %d, Type: %s\n", byteValue, byteValue, reflect.TypeOf(byteValue))

	fmt.Println("\nInferred Types:")
	fmt.Printf("autoInt: %d, Type: %s\n", autoInt, reflect.TypeOf(autoInt))
	fmt.Printf("autoFloat: %f, Type: %s\n", autoFloat, reflect.TypeOf(autoFloat))
	fmt.Printf("autoBool: %t, Type: %s\n", autoBool, reflect.TypeOf(autoBool))
	fmt.Printf("autoString: %s, Type: %s\n", autoString, reflect.TypeOf(autoString))

	fmt.Println("\nType Conversion:")
	fmt.Printf("intToFloat: %f, Type: %s\n", intToFloat, reflect.TypeOf(intToFloat))
	fmt.Printf("floatToInt: %d, Type: %s\n", floatToInt, reflect.TypeOf(floatToInt))
	fmt.Printf("intToString: %s, Type: %s\n", intToString, reflect.TypeOf(intToString))
}
