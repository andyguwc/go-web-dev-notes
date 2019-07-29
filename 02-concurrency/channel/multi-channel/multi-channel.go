/*
use select and multiple channels
use channels to signal when something is done or ready to close

select statment can watch multiple channels. Until something ahppens, it'll wait. When a channel has an event, the select statement will execute that event
*/


package main 

import (
	"fmt"
	"os"
	"time"
)

func main() {
	done := time.After(30*time.Second) // create a channel that receives a message when 30 seconds pass
	echo := make(chan []byte) // make a new channel for passing bytes from Stdin to Stdout, only hold one message at a time
	go readStdin(echo)
	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Timed Out")
			os.Exit(0)
		}
	}
}

// readSdin can only write to the out channel 
func readStdin(out chan<- []byte) { // takes a write only channel and send any received input to that channel 
	for {
		data := make([]byte, 1024)
		l,_ := os.Stdin.Read(data) // copy data from Stdin into data. Read blocks until it receives data
		if l > 0{
			out <- data // sends the buffered data over the channel 
		}
	}
}
