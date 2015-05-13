/*
 Iterate over a slice of integers, sending the odd values to an unbuffered
 channel.
 Later we drain off every odd int we saved to the channel and print its
 value.
 To show why this would be useful in the "real world", we pretend that
 some numbers have latency (perhaps it's a URI that is slow to respond).
 The "latent numbers" don't prevent the rest of the goroutines from
 finishing their work.
 This takes a little more programming effort. Unlike the last example, we
 will collect fewer return values than we input. Since channels block on
 their own, we have to keep track of the number of goroutines we create as
 well as when they complete. Otherwise the channel doesn't know when to close
 and our program will never halt.

 Based on package documentation from golang.org.
 Code is licensed under a BSD license.

*/
package main

import (
	"fmt"
	"sync" // Most definitions in this package should not be called directly
	"time"
)

// isOdd: identify odd numbers
func isOdd(i int) bool {
	// 3 and 7 shall be the numbers to simulate latency
	if i == 3 || i == 7 {
		time.Sleep(5 * time.Second)
	}

	if i%2 != 0 {
		return true
	}

	return false
}

func main() {
	// A WaitGroup counts active goroutines
	var wg sync.WaitGroup
	// unbuffered channel of integers
	odds := make(chan int)
	nums := []int{12, 2, 17, 3, 4, 5, 6, 7, 8, 9, 11} // slice literal

	for _, i := range nums { // for each nums[i]
		wg.Add(1) // increment the number of goroutines to wait for
		// goroutines also work with anonymous functions
		go func(i int) {
			// save odd values of i to the channel
			if isOdd(i) == true {
				odds <- i
			}
			defer wg.Done() // decrement the wg counter
		}(i) // call the anonymous function
	}

	// A separate goroutine waits for the other goroutines to finish
	go func() {
		wg.Wait()         // block until wg == 0 when all goroutines are finished
		defer close(odds) // program will deadlock if we don't manually close channel
	}()

	// Print the odd values we've collected
	for i := range odds {
		fmt.Println(i)
	}
}
