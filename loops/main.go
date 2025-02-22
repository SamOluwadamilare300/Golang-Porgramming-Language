package main

import "fmt"

func main(){
	x:= 0
	for x < 2 {
		fmt.Println("value of x is:", x)
		x++
	}

	for i := 0; i < 5; i++{
	fmt.Println("value is:", i)
	}

	names:= []string{"sam", "tola", "dare", "seun"}
	for i := 0; i < len(names); i++{
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf(" \n The value at index %v is %v", index, value)
	}

	for _, value := range names{
		fmt.Printf(" \n The value is %v", value)

	}
	fmt.Print("\n", names)
}