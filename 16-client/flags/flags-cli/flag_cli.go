// Examples of Flags 
// Also checkout https://gobyexample.com/command-line-flags

package main

import (
	"flag"
	"fmt"
)

// default value and short description 

// flag.String takes a flag name, default value, and description as arguments. The value of name is an address containing the value of the flag. To access this value, you’ll need to access name as a pointer
var name = flag.String("name", "World", "A name to say hello to.")

var spanish bool

func init() {
	// use existing val declared elsewhere
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language.")
}

func main() {

	// Once all flags are declared, call flag.Parse() to execute the command-line parsing
	// implicitly with long and short flags 
	flag.Parse()
	if spanish == true {
		fmt.Printf("Hola %s!\n", *name) // access as a pointer
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}

// go build flag_cli.go
// ./flag_cli -name="newname" -spanish

