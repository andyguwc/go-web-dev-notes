
// the type system implicitly resolves implemented interfaces using the methods attached to a type

// Any type that has the method String() attached, automatically implements the Stringer interface.
package main 

import (
	"fmt"
)

// Go includes the built-in interface called Stringer, defined as follows:
// type Stringer interface {
// 	String() string
// }

type metalloid struct {
	name string
	number int32
}

func (m metalloid) String() string {
	return fmt.Sprintf(
	"%-10s %-10d",
	m.name, m.number)
}

var metalloids = []metalloid{
	metalloid{"Boron", 5},
	metalloid{"Silicon", 14},
}

// The function Print(), from the standard library package fmt, will automatically call the method String(), if its parameter implements stringer

func main() {
	for _, m := range metalloids {
	fmt.Print(m, "\n")
	}
}


// interface{} type
// The following source code shows the rect type as an implementation of the interface
// type shape. The rect type is defined as a struct with receiver methods area and perim.
var shape interface {
	area() float64
	perim() float64
}

type rect struct {
	name string
	length, height float64
}

func (r *rect) area() float64 {
	return r.length * r.height
}
	
func (r *rect) perim() float64 {
	return 2*r.length + 2*r.height
}

// interface embedding 
// structure types in ways that maximize type reuse

type shape interface {
	area() float64
}

type polygon interface {
	shape
	perim()
}

type curved interface {
	shape
	circonf()
}


// Empty Interface
// all types implement the empty interfaces
// When a variable is assigned the interface{} type, the compiler relaxes its build-time type
// checks. The variable, however, still carries type information that can be queried at runtime.


// Type assertion
// Type assertion is a mechanism that is available in Go to idiomatically narrow a variable (of interface type) down to a concrete type and value that
// are stored in the variable.

// Extract the static type and value stored in the f interface parameter using assertion
// value, boolean := <interface_variable>.(concrete type name)

type food interface {
	eat()
}

type veggie string
func (v veggie) eat() {
	fmt.Println("Eating", v)
}

type meat string
func (m meat) eat() {
	fmt.Println("Eating tasty", m)
}

func eat(f food) {
	veg, ok := f.(veggie)
	if ok {
		if veg == "okra" {
			fmt.Println("Yuk! not eating ", veg)
		} else {
			veg.eat()
		}
		return 
	}
	mt, ok := f.(meat) 
	if ok {
		if mt == "beef" {
			fmt.Println("Yuk! not eating ", mt)
		} else {
			mt.eat()
		}
		return
	}

	fmt.Println("Not eating whatever that is: ", f)
}

// better way is to use swtich statement for types

func eat(f food) {
	switch mosel := f.(type) {
	case veggie:
		if morsel == "okra" {
			fmt.Println("Yuk! not eating ", mosel)
		}else {
			mosel.eat()
		}
	
	case meat:
		if morsel == "beef" {
			fmt.Println("Yuk! not eating ", mosel)
		} else {
			mosel.eat()
		}
	
	default:
		fmt.Println("Not eating whatever that is: ", f)
	}
}

