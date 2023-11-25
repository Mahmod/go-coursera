package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup

	// This function increments the shared 'counter' variable.
	increment := func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			// Race condition occurs here: multiple goroutines are accessing
			// and modifying 'counter' at the same time without synchronization.
			// This can lead to unpredictable results because the operation
			// of incrementing is not atomic.
			// It involves three steps: read the value, increment it, and write it back.
			// These steps can interleave in unexpected ways when multiple goroutines
			// are involved.
			counter++
		}
	}

	wg.Add(2)
	go increment() // Starting the first goroutine
	go increment() // Starting the second goroutine
	wg.Wait()      // Waiting for both goroutines to finish

	// The final value of 'counter' is expected to be 2000, but due to the race condition,
	// it might be less and can vary each time you run the program.
	fmt.Println("Final Counter:", counter)
}
