package main

// Struct is to what Class to other object oriented programming languages to blueprint a bill object.
//  Struct is a blueprint which describe types of data.
import "fmt"

func main(){
  myBill := newBill("Sam's expenses")
 
  myBill.addItem("pepper soup", 1000.00)
  myBill.addItem("vegetable soup", 900.00)
  myBill.addItem("egusi soup", 700.00)
  myBill.addItem("ewedu soup", 800.00)

  myBill.updateTip(10)

  fmt.Println(myBill.format())
}


