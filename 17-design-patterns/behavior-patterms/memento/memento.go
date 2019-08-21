
/* 
Remember actions with this design pattern 

Memento is all about a sequence of actions over time, say to undo one or two operations or
to provide some kind of transactionality to some application.


Memento provides the foundations for many tasks, but its main objectives could be defined as follows:
- Capture an object state without modifying the object itself
- Save a limited amount of states so we can retrieve them later
*/


package memento

import "fmt"


// Memento design pattern is usually composed of three actors:
// state, memento, and originator

type memento struct {
	state State
}

type State struct {
	Description string
}

//--------------------------------------------------------------------

type originator struct {
	state State
}

// The NewMemento method will return a new Memento built with originator current State value
func (o *originator) NewMemento() memento {
	return memento{state: o.state}
}

// The ExtractAndStoreState method will take the state of a Memento and store it in the Originator's state field
func (o *originator) ExtractAndStoreState(m memento) {
	o.state = m.state
}


//--------------------------------------------------------------------

type careTaker struct {
	mementoList []memento
}

func (c *careTaker) Add(m memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) Memento(i int) (memento, error) {
	if len(c.mementoList) < i || i < 0 {
		return memento{}, fmt.Errorf("Index not found\n")
	}
	return c.mementoList[i], nil
}

