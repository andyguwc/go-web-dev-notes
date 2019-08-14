/* use waitgroup 

use case: each goroutine to compress the file then main program waits until all done 

A wait group is a message-passing facility that signals a waiting goroutine when it’s
safe to proceed. To use it, you tell the wait group when you want it to wait for something,
and then you signal it again when that thing is done. A wait group doesn’t need
to know more about the things it’s waiting for other than (a) the number of things it’s
waiting for, and (b) when each thing is done. You increment the first with wg.Add,
and as your task completes, you signal this with wg.Done. The wg.Wait function blocks
until all tasks that were added are done. 
*/

package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)



func main() {
	var wg sync.WaitGroup // wait group, no need to initialize
	var i int
	var file string

	for i, file = range os.Args[1:] { // collect list of files passed in on the command line 
		wg.Add(1) // for every file you add, tell the wait group that you're waiting for one more compress operation
		go func(filename string) {
			compress(filename) // function call compresses then notifies wait group it's done
			wg.Done()
		}(file) // passing in file parameter
	}
	wg.Wait() // outer goroutine waits until all the compressing goroutines have called wg.Done

	fmt.Printf("Compressed %d files\n", i+1)
}

func compress(filename string) error {
	fmt.Printf("Compressing %s\n", filename)
	in, err := os.Open(filename) // open source file for reading 
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz") // open a destination file for writing to 
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}

// go run compress.go exampledata/*

