//Golang is a generic pass-by-value language
//-Go makes copies of values when passed into functions.
package main

import "fmt"

func updateName(x string) string{
	
   x = "Dare"
   return x 
}

func updateMenu(y map[string]float64){
	y["coffee"] = 200.67
}

func main(){
	//first group type: strings, ints, bools, floats, structs
	name := "Sam"

	name = updateName(name) 
	fmt.Println(name)
	//second group type: slices, maps, functions
	menu := map[string]float64{
	"pie":          100.00,
	 "ice cream":       300.00,
	}

	updateMenu(menu)
	fmt.Println(menu)
}