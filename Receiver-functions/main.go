package main

// Struct is to what Class to other object oriented programming languages to blueprint a bill object.
//  Struct is a blueprint which describe types of data.
import "fmt"

func main(){
  myBill := newBill("Sam's expenses")
 

  fmt.Println(myBill.format())
}


