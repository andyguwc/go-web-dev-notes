/* 
The objective of the chain of responsibility is to provide to the developer a way to chain
actions at runtime. The actions are chained to each other and each link will execute some
action and pass the request to the next link (or not).

Example: multi logger chain 
*/


package chain_of_responsability

import (
	"fmt"
	"io"
	"strings"
)

// a chain of responsibility interface will have a Next() method. The Next() method executes the next link in the chain
type ChainLogger interface {
	Next(string)
}

// simple logger that logs the text of a request 

// The FirstLogger and SecondLogger types have exactly the same structure--both
// implement ChainLogger and have a NextChain field that points to the next ChainLogger
type FirstLogger struct {
	NextChain ChainLogger
}

func (f *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)

	if f.NextChain != nil {
		f.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

// if incoming text has "hello" then write on the console
func (se *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n", s)

		if se.NextChain != nil {
			se.NextChain.Next(s)
		}

		return
	}

	fmt.Printf("Finishing in second logging\n\n")
}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (w *WriterLogger) Next(s string) {
	if w.Writer != nil {
		w.Writer.Write([]byte("WriterLogger: " + s))
	}

	if w.NextChain != nil {
		w.NextChain.Next(s)
	}
}

// Closure is a function that takes a string and returns nothing

type ClosureChain struct {
	NextChain ChainLogger
	Closure   func(string)
}

func (c *ClosureChain) Next(s string) {
	if c.Closure != nil {
		c.Closure(s)
	}

	if c.NextChain != nil {
		c.Next(s)
	}
}
