package main

import "fmt"

func main() {
	age := 20
	Age := "30"
	name := "Samuel Afolabi"
	occupation := "Software Developer"

	fmt.Println("Hello,")
	fmt.Println("World!")

	fmt.Println("My name is", name, ",I am a", occupation, ",and I am", age, "years old")
	fmt.Println("I am a", occupation)
	fmt.Println("I am", age, "years old")

	///Formatted string &_ = format specifier
	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", Age, name)
	fmt.Printf("age is of type %T and %T  and name is of type %T \n", age, Age, name)
	fmt.Printf("you scored %0.f points \n", 9.0)

	//Sprintf (Save formatted string)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is:", str)

}

//Learn More on Format Specifiers
// %v - default format for variables
// %T - data type of the variable
// %q - double-quoted strings
// %d - base 10
// %b - base 2
// %o - base 8
// %x - base 16
// %c - character
// %s - plain string
// %p - pointer
// %f - float
// %e - scientific notation
// %t - boolean
// %U - unicode
// %x - base 16, with lower-case letters for a-f
// %X - base 16, with upper-case letters for A-F
// %v - default format for variables
// %T - data type of the variable
// %q - double-quoted strings
// %d - base 10
// %b - base 2
// %o - base 8
// %x - base 16
// %c - character
// %s - plain string
//
