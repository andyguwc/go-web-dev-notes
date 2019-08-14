// Pointers

// In Go, when a piece of data is stored in memory, the value for that data may be accessed
// directly or a pointer may be used to reference the memory address where the data is located.

// Pointer type

package main
import "fmt"
	var valPtr *float32
	var countPtr *int
	var person *struct {
		name string
		age int
	}
	var matrix *[1024]int
	var row []*int64
	func main() {
		fmt.Println(valPtr, countPtr, person, matrix, row)
}

// Address operator
var a int = 1024
var aptr *int = &a

// The assigned address value will always be the same (always pointing to a) regardless of
// where aptr may be accessed in the code


// new() function returns a pointer 

// It first allocates the appropriate memory for a zero-value of the specified type. The function then
// returns the address for the newly created value.

func main() {
	intptr := new(int)
	*intptr = 44

	p := new(struct{ first, last string })
	// It is not necessary to write *p.first to access the pointer's field value. We can drop the * and just use p.first
	p.first = "Samuel"
	p.last = "Pierre"
	fmt.Printf("Value %d, type %T\n", *intptr, intptr)
	fmt.Printf("Person %+v\n", p)
}

