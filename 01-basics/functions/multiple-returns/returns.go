package main

import (
	"fmt"
)

// return two values
func Names() (string, string) {
	return "Foo", "Bar"
}

// returned values have names
// values assigned to returne values and call return 
func Names2() (first string, second string) {
	first = "Foo"
	second = "Bar"
	return
}

func main() {
	n1, n2 := Names()
	fmt.Println(n1,n2)

	// skip on the second value
	n3, _ := Names2()
	fmt.Println(n3)
}

