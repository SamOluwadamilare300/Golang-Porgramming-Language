
package main

import "fmt"

//Map allow us to store key values pairs.
// The key and can be different type but in SINGLE and all the values must be the same. 
func main(){
	menu := map[string]float64{
		"eba": 200.00,
		"fufu": 400.00,
		"rice": 300.00,
		"salad": 500.00, 
	}
	fmt.Println(menu)
	fmt.Println(menu["eba"])

	//Looping Through Maps
   for k, v := range menu{
	fmt.Println(k, "-", v)
   }

   // Ints as Key Type

   phoneBook := map[int]string{
    35557373: "sam",
	43526388: "dare",
	54263883:  "folabi",
   }
   fmt.Println(phoneBook)
   fmt.Println(phoneBook[54263883])
   phoneBook[35557373] = "sammy"
   fmt.Println(phoneBook)
   phoneBook[43526388] = "dammy"
   fmt.Println(phoneBook)
}