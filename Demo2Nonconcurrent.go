/*
 Iterate over a slice of integers, saving the odd values.
 To show why this would be useful in the "real world", we pretend that
 some numbers have latency (perhaps it's a URI that is slow to respond).
 This version takes twice as long to run because the "latent" values
 stall the program for every call to isOdd().

 Based on package documentation from golang.org.
 Code is licensed under a BSD license.

*/
package main

import (
	"fmt"
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
	odds := make([]int, 0)                            // create an empty slice
	nums := []int{12, 2, 17, 3, 4, 5, 6, 7, 8, 9, 11} // slice literal

	for _, i := range nums { // for each nums[i]
		// save odd values of i to the slice
		if isOdd(i) == true {
			odds = append(odds, i)
		}
	}

	// Print the odd values we've collected
	for _, i := range odds {
		fmt.Println(i)
	}
}
