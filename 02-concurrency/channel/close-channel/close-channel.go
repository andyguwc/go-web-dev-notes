/*
Safely close channels and exit goroutines 
use a channel (often called done) to send a signal 
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	done := make(chan bool) // indicates when you are finished 
	until := time.After(5 * time.Second)

	go send(msg, done) // passes two channels into send 

	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			done <- true // when you time out, let send know the prcess is done
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

func send(ch chan<- string, done <-chan bool) { // ch a receiving channel while done a sending channel 
	for {
		select {
		case <-done:
			println("Done")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}