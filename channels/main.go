package main

import (
	"fmt"
	"sync"
)

// worker receives a value from the 'in' channel, increments it,
// sends the result to the 'out' channel, and then closes 'out'.
// The WaitGroup pointer is passed explicitly.
func worker(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Read a value from the input channel.
	// The second value (ok) tells us if the channel is still open.
	if val, ok := <-in; ok {
		out <- val + 1
	}
	close(out) // Signal that no more values will be sent on out.
}

// numberGenerator generates sequential numbers from 0 up to (but not including) max,
// sends them on the 'out' channel, and then closes 'out'.
// The WaitGroup pointer is passed explicitly.
func numberGenerator(max int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out) // Close the channel when done.

	for i := 0; i < max; i++ {
		out <- i
	}
}

func main() {
	var wg sync.WaitGroup

	// Create channels.
	// Using a buffered channel for inChan prevents any potential blocking when sending.
	inChan := make(chan int, 1)
	outChan := make(chan int)
	numberChan := make(chan int)

	// Start the worker goroutine.
	wg.Add(1)
	go worker(inChan, outChan, &wg)

	// Send the initial value to the worker and then close the channel
	// because no further values will be sent.
	inChan <- 5
	close(inChan)

	// Receive the result from the worker.
	// Since the worker closes outChan when done, a single read is sufficient.
	result, ok := <-outChan
	if !ok {
		fmt.Println("Worker output channel closed unexpectedly")
		return
	}
	fmt.Printf("Worker processed value: %d\n", result)

	// Start the number generator with the result from the worker.
	wg.Add(1)
	go numberGenerator(result, numberChan, &wg)

	// Print generated numbers.
	// The loop will exit when numberChan is closed by numberGenerator.
	for num := range numberChan {
		fmt.Printf("Generated number: %d\n", num)
	}

	// Wait for all goroutines to finish.
	wg.Wait()
}
