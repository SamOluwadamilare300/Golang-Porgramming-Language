package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	options, _ := getInput("Choose options (A - Add Book, S - Save the Bill, T - Add Tip):", reader)

	switch options {
	case "A":
		name, _ := getInput("Book name:", reader)
		price, _ := getInput("Book Price:", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Book added -", name, price)
		promptOptions(b)

	case "T":
		tip, _ := getInput("Enter tip amount (NGN):", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip added -", tip)
		promptOptions(b)

	case "S":
		b.Save()
		fmt.Println("You saved the file ✅", b.name)

	default:
		fmt.Println("That was not a valid option! ❌")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}