/* Buffered Channel
If a channel has received a value, and then sent another one before the channel can be read, the second operation will block
so we need a buffered channel 


Using channel as a lock
1 A function acquires a lock by sending a message on a channel.
2 The function proceeds to do its sensitive operations.
3 The function releases the lock by reading the message back off the channel.
4 Any function that tries to acquire the lock before it’s been released will pause
when it tries to acquire the (already locked) lock.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// create a buffered channel with one space 
	lock := make(chan bool, 1)

	// starts up goroutines sharing the locking channel 
	for i := 1; i < 7; i++ {
		go worker(i, lock)
	}
	time.Sleep(10 * time.Second)
}

func worker(id int, lock chan bool) {
	fmt.Printf("%d wants the lock\n", id)

	// A worker acquires the lock by sending it a message. The first worker to hit this will get the one space, and thus own the lock. The rest will block.
	lock <- true
	fmt.Printf("%d has the lock\n", id) // this section is locked 
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d is releasing the lock\n", id)
	<-lock // releasing the lock by reading the value, which then opens that one space on the buffer again 
}
