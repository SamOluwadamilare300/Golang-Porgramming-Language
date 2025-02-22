package main

import (
	"fmt"
	"time"
)

func main() {
	// String variables
	var nameOne string = "Sam"
	var nameTwo = "Dare"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree) // Should print: Sam Dare Placeholder

	// Override variables
	nameOne = "Samuel"
	nameThree = "Afolabi"

	fmt.Println(nameOne, nameTwo, nameThree)
	nameFour := "Samson Afolabi"
	fmt.Println(nameFour)

	// Integer variables
	var oldOne int = 20
	var oldtwo = 40
	oldThree := 60

	fmt.Println(oldOne, oldtwo, oldThree)

	//Integer Format Specifiers
	fmt.Printf("Decimal: %d, Binary: %b, Octal: %o, Hex: %x\n", 255, 255, 255, 255)
	// Output: Decimal: 255, Binary: 11111111, Octal: 377, Hex: ff

	// Floating-Point Format Specifiers
	fmt.Printf("Float: %f, Scientific: %e\n", 123.456, 123.456)
	// Output: Float: 123.456000, Scientific: 1.234560e+02

	// String Format Specifiers
	fmt.Printf("String: %s, Quoted: %q, Hex: %x\n", "hello", "hello", "hello")
	// Output: String: hello, Quoted: "hello", Hex: 68656c6c6f

	// Boolean Format Specifier
	fmt.Printf("Boolean: %t\n", true)
	// Output: Boolean: true

	//Pointer Format Specifier: %p	Pointer address in hexadecimal
	var num = 42
	fmt.Printf("Pointer: %p\n", &num)
	// Output: Pointer: 0xc0000100b0 (actual address varies)

	// Set a specific date and time: (Year, Month, Day, Hour, Minute, Second, Nanosecond, Timezone)
	t := time.Date(2025, time.February, 21, 14, 30, 45, 0, time.UTC)

	// Print formatted time
	fmt.Printf("Current Time: %02d-%02d-%d %02d:%02d:%02d\n",
		t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())

}
