/* Command Pattern

When using the Command design pattern, we are trying to encapsulate some sort of action
or information in a light package that must be processed somewhere else.
- Put some information into a box. Just the receiver will open the box and know its
contents.
- Delegate some action somewhere else.


We need a constructor of console printing commands. When using this
constructor with a string, it will return a command that will print it. In this case,
the handler is inside the command that acts as a box and as a handler.

We need a data structure that stores incoming commands in a queue and prints
them once the queue reaches the length of three.

*/


package main

import (
	"fmt"
	"net/http"
)

type Command interface {
	Execute()
}

// need some type implementing this interface 
type ConsoleOutput struct {
	message string
}

// implements the Command interface and prints to the console 
func (c *ConsoleOutput) Execute() {
	fmt.Println(c.message)
}

// a Command constructor that accepts a message string and returns the Command interface 
func CreateCommand(s string) Command {
	fmt.Println("Creating command")

	return &ConsoleOutput{
		message: s,
	}
}

// stores an array of the Commands interface 
type CommandQueue struct {
	queue []Command
}

func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)

	if len(p.queue) == 3 {
		for _, command := range p.queue {
			command.Execute()
		}

		p.queue = make([]Command, 3)
	}
}

func main() {
	queue := CommandQueue{}

	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))

	queue.AddCommand(CreateCommand("Fourth message"))
	queue.AddCommand(CreateCommand("Fifth message"))

	client := http.Client{}
	client.Do(nil)
}