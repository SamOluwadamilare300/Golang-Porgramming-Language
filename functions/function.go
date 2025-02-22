package main

import (
	"fmt"
	"math"
)

//Functions
func sayHello(n string){
	fmt.Printf("Good Morning %v \n", n)
}

func sayPraise(n string){
	fmt.Printf("God is Good %v \n", n)
}
func cycleValues(n []string, f func(string)){
	for _, v := range n {
		f(v)
	}
}

func sayGreetings(n string){
	fmt.Printf("GoodBye %v \n", n)
}

func circleArea(r float64)  float64{
	return math.Pi * r * r 
}

func main(){
	sayHello("Sammy")
	sayPraise("All the time")
    sayGreetings("Damilare")
	cycleValues([]string{"Love", "Friendship", "Sleep"}, sayGreetings)
	cycleValues([]string{"Today", "Yestarday", "Forever"}, sayPraise)
     a1 := circleArea(10.5)
	 a2 := circleArea(15)

	fmt.Println(a1,a2)
	fmt.Printf("Circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)
}