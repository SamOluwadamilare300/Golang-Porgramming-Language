package main

import "fmt";


func updateName(x *string){
	*x = "Damilare"
}

func main(){
	name := "Sammy"

	// updateName(name)

	fmt.Println("memory address of name is:", &name)
	
	m :=&name
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address", *m)

	fmt.Println(name)
    fmt.Println(name)
	updateName(m)
	fmt.Println(name)
}