package main

import (
	"fmt"
	"sync"
	"time"
)

func processTransaction(transactionNumber int, done chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Processing transaction %d\n", transactionNumber)
	time.Sleep(2 * time.Second)
	done <- transactionNumber
}

func main() {
	totalTransaction := 5
	processed := make(chan int, totalTransaction)
	var wg sync.WaitGroup

	for i := 1; i <= totalTransaction; i++ {
		wg.Add(1)
		go processTransaction(i, processed, &wg)
	}

	go func() {
		wg.Wait()
		close(processed)
	}()

	for transactionNumber := range processed {
		fmt.Printf("Received completion signal for transaction %d\n", transactionNumber)
	}

	fmt.Println("All transactions processed!")
}