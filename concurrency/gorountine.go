package main
//The Gorountine create threads to execute the code concurrently
//  without an advance knowledge of any particular.
import (
	"fmt"
	"time"
)

func processTransaction(transactionNumber int){
	fmt.Println("Process transaction", transactionNumber)
	time.Sleep( 2 * time.Second) // simulate time-consuming task.
	fmt.Println("Process transaction", transactionNumber)
}

func main(){
  for i := 1; i<= 5; i++{
	go processTransaction(i)  //start a gorountine for each transaction
  }

  time.Sleep(3 * time.Second)   // wait for all gorountine to finish
  fmt.Println("All transactions processed!")
}

