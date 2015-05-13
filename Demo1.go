/*
  Demonstrate how to use channels and goroutines to keep the program alive.
  Iterate over a slice of numbers, passing each each i to a function
  that calculates the base10 log of i.
  We don't actually care what the return value is.
  Instead, the function just signals a channel that its work is done.

  Based on package documentation from golang.org.
  Code is licensed under a BSD license.
*/
package main

import (
	"fmt"
	"math/cmplx"
)

// print the decimal log of a number and signal completion
func getLog(c chan bool, i complex128) {
	fmt.Printf("%v Log: %v\n", i, cmplx.Log10(i)) // %v for any value
	c <- true                                     // Send signal to channel
}

func main() {
	// buffered channel of bool; doesn't need a receiver
	// this effectively makes the done channel a 'first in, first out' queue
	done := make(chan bool, 1)
	nums := []complex128{7, 8, cmplx.Sqrt(-9), 10} // slice literal

	for _, i := range nums { // _ = 0,1, ...; i = 7, 8, ...
		go getLog(done, i) // run getLog() as goroutine; don't wait for return
	}

	<-done // empty the 'done' channel, discarding its value
}
