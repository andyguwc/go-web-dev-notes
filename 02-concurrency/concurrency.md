# Concurrency

In concurrency, the tasks don’t necessarily need to start or end together—their execution overlaps. These tasks are scheduled and often (though not necessarily) communicate to share data as well as coordinate the execution times.

In parallelism, tasks start and are executed at the same time. Usually a larger problem
is split into smaller chunks and processed simultaneously to improve performance.
Parallelism usually requires independent resources (for example, CPUs);
concurrency uses and shares the same resources.


# Goroutines
Goroutines are functions that run independently with other goroutines. This might
seem similar to threads—and in fact, goroutines are multiplexed on threads—but
they aren’t threads. A lot more goroutines than threads can be running, because
goroutines are lightweight. A goroutine starts with a small stack (8 K as of Go 1.4) and
it can grow (or shrink) as needed. Whenever a goroutine is blocked, it blocks the OS
thread it’s multiplexed on, but the runtime moves other goroutines on the same
blocked thread to another unblocked thread.


package main 

func printNumbers2() {
    for i := 0; i < 10; i++ {
        time.Sleep(1 * time.Microsecond
        fmt.Printf("%d ", i)
    }
}
func printLetters2() {
    for i := 'A'; i < 'A'+10; i++ {
        time.Sleep(1 * time.Microsecond)
        fmt.Printf("%c ", i)
    }
}
func goPrint2() {
    go printNumbers2()
    go printLetters2()
}


## Benchmark 
The benchmarks are worse because of the same issue I mentioned earlier: scheduling
and running on multiple CPUs have a cost, and if the processing doesn’t warrant the
high cost, it can make the performance worse.

## Waiting for Goroutines
Ensure all go routines complete before moving to the next thing 
The mechanism is straightforward:
- Declare a WaitGroup.
- Set up the WaitGroup’s counter using the Add method.
- Decrement the counter using the Done method whenever a goroutine completes
its task.
- Call the Wait method, which will block until the counter is 0.

```
package main

import "fmt"
import "time"
import "sync"

func printNumbers2(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers2(&wg)
	go printLetters2(&wg)
	wg.Wait()
}
```

# Channels
Channels are typed values that allow goroutines to communicate with each other.
Channels are allocated using make, and the resulting value is a reference to an underlying
data structure. This, for example, allocates a channel of integers:
ch := make(chan int)

Channels are, by default, unbuffered. If an optional integer parameter is provided, a
buffered channel of the given size is allocated instead. This creates a buffered channel of
integers with the size 10:
ch := make(chan int, 10)

Unbuffered channels are synchronous. You can think of unbuffered channels as a box
that can contain only one thing at a time. Once a goroutine puts something into this
box, no other goroutines can put anything in, unless another goroutine takes out
whatever is inside it first. This means if another goroutine wants to put in something
else when the box contains something already, it will block and go to sleep until the
box is empty.


This allocates a send-only channel of strings:
ch := make(chan <- string)

This allocates a receive-only channel of strings:
ch := make(<-chan string)


## Buffered Channels
c := make(chan int, 3)


## Selecting Channels 
Select is like a switch statement for channels 

## Closing Channels
Channels can be closed using the close built-in function. Closing a channel indicates
to the receiver that no more values will be sent to the channel. You can’t close a 
receive-only channel, and sending to or closing an already closed channel causes a
panic. 

A closed channel is never blocked and always returns the zero value for the
channel’s type.

# Race Conditions
A race condition exists when the program depends on a specific sequence or timing
for it to happen and specific sequence or timing can’t be guaranteed. As a result, the
behavior of the program becomes erratic and unpredictable.

Race conditions commonly appear in concurrent programs that modify a shared resource.
If two or more processes or threads try to modify the shared resource at the
same time, the one that gets to the resource first will behave as expected but the
other processes won’t. Because we can’t predict which process gets the resource
first, the system won’t behave consistently.

