package main

import "fmt"

func worker(ch chan int) {
	val := <-ch   // This process is blocked until the channel receives a value
	ch <- val * 2 // Send a processed value back into the channel
}

func main() {
	ch := make(chan int)

	go worker(ch) // Start the worker goroutine
	//ch <- 42      // Send a value to the channel
	//fmt.Println(<-ch) // Receive the processed value from the channel
	ch <- 42          // Send a value to the channel
	fmt.Println(<-ch) // Receive the processed value from the channel
}
