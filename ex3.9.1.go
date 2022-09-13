package main

import (
	"fmt"
)

func main() {
	fmt.Println("Started!")

	calculator := func(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
		result := make(chan int)

		go func() {
			defer close(result)

			select {
			case value := <-firstChan:
				result <- value * value
			case value := <-secondChan:
				result <- value * 3
			case <-stopChan:
				return
			}
		}()

		return result
	}

	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	go func() {
		// NOTE: для повторного использования канала лучше не закрывать канал, а отправить в него
		// done <- struct{}{}
		defer close(stopChan)

		firstChan <- 10
		firstChan <- 100
		secondChan <- 10
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("Result:", <-calculator(firstChan, secondChan, stopChan))
	}
	// Result: 100
	// Result: 10000
	// Result: 30
	// Result: 0

	fmt.Println("Finish!")
}
