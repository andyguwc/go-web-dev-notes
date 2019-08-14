/* 

WaitGroup is often used as a way to implement work distribution patterns. The following
code snippet illustrates work distribution to calculate the sum of multiples of 3 and 5 up to
MAX. The code uses the WaitGroup variable, wg, to create a concurrency barrier that waits
for two goroutines to calculate the partial sums of the numbers, then gathers the result after
all goroutines are done

*/

package main

import (
	"fmt"
	"sync"
)

const MAX = 1000

// The method call, wg.Add(2), configures the WaitGroup variable wg
// because the work is distributed between two goroutines.

func main() {
	values := make(chan int, MAX)
	result := make(chan int, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() { // gen multiple of 3 & 5 values
		for i := 1; i < MAX; i++ {
			if (i%3) == 0 || (i%5) == 0 {
				values <- i // push downstream
			}
		}
		close(values)
	}()

	work := func() { // work unit, calc partial result
		// wg.Done() to decrement the WaitGroup counter by one every time it is completed.
		defer wg.Done()
		r := 0
		for i := range values {
			r += i
		}
		result <- r
	}

	// distribute work to two goroutines
	go work()
	go work()

	wg.Wait()                    // wait for both groutines
	total := <-result + <-result // gather partial results
	fmt.Println("Total:", total)
}
