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
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

//Format The Bill

func (b *bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0

	//List Items

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...NGN%v \n", k+":", v)

		total += v
	}

	//add tip
	fs += fmt.Sprintf("%-25v ...NGN%v\n", "tip:", b.tip)

	//Get total
	fs += fmt.Sprintf("%-25v ...NGN%0.2f", "total:", total+ b.tip)

	return fs
}

//Updating tips

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

//Adding Items to the Bill

func (b bill) addItem(name string, price float64) {
	b.items[name] = price
}
