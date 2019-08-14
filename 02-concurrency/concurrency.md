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

Goroutines can also be defined as function ligerals 
func main() {
	go count(10, 30, 10)
	go func() {
		count(40, 60, 10)
	}()
	...
}


If these values are updated outside of the closure, it may create
race conditions causing the goroutine to read unexpected values by the time it is scheduled
to run.

The following snippet shows an example where the goroutine closure captures the variable j from the loop:
func main() {
	starts := []int{10,40,70,100}
	for _, j := range starts{
		go func() {
			count(j, j+20, 10)
		}()
	}
}

Since j is updated with each iteration, it is impossible to determine what value will be read by the closure. In most cases, the goroutine closures will see the last updated value of j by the time they are executed. This can be easily fixed by passing the variable as a parameter in the function literal for the goroutine, as shown here

func main() {
	starts := []int{10,40,70,100}
	for _, j := range starts{
		go func(s int) {
			count(s, s+20, 10)
		}(j)
	}
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
## Goroutine Runtime Scheduling

Go's runtime scheduler uses a form of cooperative scheduling to schedule goroutines. 

By default, the scheduler will allow a running goroutine to execute to completion. However, the scheduler will automatically yield to another goroutine for execution if one of the following events occurs:
- A go statement is encountered in the executing goroutine
- A channel operation is encountered (channels are covered later)
- A blocking system call (file or network IO for instance) is encountered
- After the completion of a garbage collection cycle

The scheduler will schedule a queued goroutines ready to enter execution when one of the previous events is encountered in a running goroutine. It is important to point out that the scheduler makes no guarantee of the order of execution of goroutines. 

When the following code snippet is executed, for instance, the output will be printed in an arbitrary order for each run.


# Channels
Channels are typed values that allow goroutines to communicate with each other.
Channels are allocated using make, and the resulting value is a reference to an underlying data structure. This, for example, allocates a channel of integers:
ch := make(chan int)

Channels are, by default, unbuffered. If an optional integer parameter is provided, a
buffered channel of the given size is allocated instead. This creates a buffered channel of
integers with the size 10:
ch := make(chan int, 10)


## Unbuffered Channel
Unbuffered channels are synchronous. You can think of unbuffered channels as a box
that can contain only one thing at a time. Once a goroutine puts something into this
box, no other goroutines can put anything in, unless another goroutine takes out
whatever is inside it first. This means if another goroutine wants to put in something else when the box contains something already, it will block and go to sleep until the box is empty.

ch := make(chan int) // unbuffered channel
- If the channel is empty, the receiver blocks until there is data
- The sender can send only to an empty channel and blocks until the next receive
operation
- When the channel has data, the receiver can proceed to receive the data.

Recall that the sender blocks immediately upon sending to an unbuffered channel. This
means any subsequent statement, to receive from the channel for instance, becomes
unreachable, causing a deadlock. The following code shows the proper way to send to an unbuffered channel:

func main() {
	ch := make(chan int)
	go func() { ch <- 12 }()
	fmt.Println(<-ch)
}


## Buffered Channels
c := make(chan int, 3)

- When the channel is empty, the receiver blocks until there is at least one element
- The sender always succeeds as long as the channel is not at capacity
- When the channel is at capacity, the sender blocks until at least one element is received

## Unidirectional Channels

This allocates a send-only channel of strings:
ch := make(chan <- string)

This allocates a receive-only channel of strings:
ch := make(<-chan string)

// pings is a channle for receives
// pongs is a channel for sends 
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

## Selecting Channels 
Select is like a switch statement for channels 


## Closing Channels
Channels can be closed using the close built-in function. Closing a channel indicates
to the receiver that no more values will be sent to the channel. You can’t close a 
receive-only channel, and sending to or closing an already closed channel causes a
panic. 

A closed channel is never blocked and always returns the zero value for the
channel’s type.


// Example implementing a word histogram. Reads the words from the data slice and on a separate goroutine collects the occurrence of each word

func main() {
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink ...",
		"The dark bird bird of prey lands on a small ...",
	}

	histogram := make(map[string]int)

	done := make(chan bool)

	//splits and count words
	go func() {
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				histogram[word]++
			}
		}
	done <- true
	}()

	if <-done {
		for k, v := range histogram {
			fmt.Printf("%s\t(%d)\n", k, v)
		}
	}

## Streaming Data
Stream data from one goroutine to another.
- Continuously send data on a channel
- Continuously receive data from that channel
- Signal the end of the stream so receiver may stop


## Generator Functions
In this approach, a goroutine is wrapped in a function which generates values that are sent via a channel returned by the function. The consumer goroutine receives these values as they are generated.

// generator function that produces data
func words(data []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out) // closes channel upon fn return
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				out <- word
			}
		}
	}()
	return out
}


## Select from Multiple Channels
select {
	case <send_ or_receive_expression>:
	default:
}

# Wait Group 
Sometimes when working with goroutines, you may need to create a synchronization
barrier where you wish to wait for all running goroutines to finish before proceeding. The sync.WaitGroup type is designed for such a scenario, allowing multiple goroutines torendezvous at specific point in the code. Using WaitGroup requires three things:
- The number of participants in the group via the Add method
- Each goroutine calls the Done method to signal completion
- Use the Wait method to block until all goroutines are done



# Race Conditions
A race condition exists when the program depends on a specific sequence or timing
for it to happen and specific sequence or timing can’t be guaranteed. As a result, the
behavior of the program becomes erratic and unpredictable.

Race conditions commonly appear in concurrent programs that modify a shared resource.
If two or more processes or threads try to modify the shared resource at the
same time, the one that gets to the resource first will behave as expected but the
other processes won’t. Because we can’t predict which process gets the resource
first, the system won’t behave consistently.

## Debugging Race Conditions

The compiler's output shows the offending goroutine locations that caused the race condition

go run -race sync1.go





