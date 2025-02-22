package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//Another bill

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 500.00, "cake": 400.99, "bread": 300.00},
		tip:   0,
	}
	return b
}

//Format The Bill

func (b bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0

	//List Items

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...NGN%v \n", k+":", v)

		total += v
	}

	//Get total
	fs += fmt.Sprintf("%-25v ...NGN%0.2f", "total:", total)

	return fs
}