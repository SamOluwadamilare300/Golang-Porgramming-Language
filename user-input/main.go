package main

// Struct is to what Class to other object oriented programming languages to blueprint a bill object.
//  Struct is a blueprint which describe types of data.
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name:", reader)

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	options, _ := getInput("Choose options (A -add item, S - save the bill, T - add tip):", reader)
     
	switch options {
	case "A":
		name, _ := getInput("Book name:", reader)
		price, _ :=getInput("Book Price:", reader)
		fmt.Println(name, price)
		
	case "T":
		tip, _ :=getInput("Enter tip amount (NGN):", reader)
		fmt.Println(tip)

	case "S":
		fmt.Println("You chose S ✅")	
	default: 
	  fmt.Println("That wa not a Valid Options!❌")
	  promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)

}
