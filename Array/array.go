
package main

import "fmt"


func main(){

	///The lenght of array can not change but the lenght of slices can change. 
	// and we get ranges from existing array and store them in a slices.
	
	var ages[4] int = [4]int{20, 25, 30, 45}

	names := [4]string{"John", "Doe", "Smith", "Alex"}
	names[1]= "Samson";

	fmt.Println(ages, len(ages))
	fmt.Println(ages[1])
	fmt.Println(names, len(names))
	fmt.Println(names[2])
	

	var scores = []int{10, 20, 30, 40, 50}
	scores[2] = 23
	scores = append(scores, 85)

	fmt.Println(scores, len(scores));
    
	//Slice Range
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree);

	rangeOne = append(rangeOne, "Damilare")
	fmt.Println(rangeOne)

}