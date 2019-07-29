/* Custom Error
implement the error interface 
have additional information on where in the file the error occured 

type error interface {
	Error() string
}
*/


package main

import "fmt"

func main() {

	err := &ParseError{
		Message: "Unexpected char ';'",
		Line:    5,
		Char:    38,
	}

	fmt.Println(err.Error())
}

type ParseError struct {
	Message    string
	Line, Char int // the location information 
}

func (p *ParseError) Error() string { // implement the error interface 
	format := "%s on Line %d, Char %d"
	return fmt.Sprintf(format, p.Message, p.Line, p.Char)
}
