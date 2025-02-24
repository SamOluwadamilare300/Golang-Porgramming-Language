package main
//The wait group in concurrency
import (
	"fmt"
	"sync"
	"time"
)

func main(){

	var wg sync.WaitGroup
  for i := 1; i<= 5; i++{
	wg.Add(1)
	go func(transactionNumber int){
		defer wg.Done()
		fmt.Println("Processing transaction", transactionNumber)
		time.Sleep(2 * time.Second)   // wait for all gorountine to finish
		fmt.Println("All transactions processed!")
	}(i)
  }

  wg.Wait()
  fmt.Println("All transactions processed!")
}

