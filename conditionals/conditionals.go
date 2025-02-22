//Booleans and Conditionals 

package main

import "fmt"

func main(){
	age := 30
	fmt.Println(age <= 70)
	fmt.Println(age >= 70)
	fmt.Println(age == 55)
	fmt.Println(age != 70)

	if age < 30 {
		fmt.Println("age is less than 30")
	}else if age < 55 {
		fmt.Println("age is less than 55")
	}else {
		fmt.Println("age is not less than 60")
	}

  names := []string{"sam", "dare","tola","deji","folabi"}
  for index, value := range names{
	if index == 1 {
		fmt.Println("continuing at position", index)
		continue
	}
	if index > 2 {
		fmt.Println("Breaking at position", index)
		break
	}
	fmt.Printf("The value at position %v is %v \n", index, value)
  }
}