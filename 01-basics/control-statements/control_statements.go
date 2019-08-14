
// if statement

if {

} else if {

} else {

}


if num0 > 100 || num0 < 900 {
	fmt.Println("Currency: ", num0)
	printCurr(num0)
}

// also supports intialization executed before the test expression

if num1 := 300; num1 > 100 || num1 <500 {
	fmt.Println(num1)
}

// switch statement
switch i {
case 1:
	fmt.Println("one")
case 2:
	fmt.Println("two")
case 3:
	fmt.Println("three")
}

// default for switch 
switch time.Now().Weekday() {
case time.Saturday, time.Sunday:
	fmt.Println("It's the weekend")
default:
	fmt.Println("It's a weekday")
}

// fallthrough cases 
// When a case is matched, the fallthrough statements cascade down to the first statement of the successive case block.

// expressionless switches
// Go supports a form of the switch statement that does not specify an expression. In this
// format, each case expression must evaluate to a Boolean value true.

import (
	"fmt"
	"strings"
)
type Curr struct {
	Currency string
	Name string
	Country string
	Number int
}
var currencies = []Curr{
	Curr{"DZD", "Algerian Dinar", "Algeria", 12},
	Curr{"AUD", "Australian Dollar", "Australia", 36},
	Curr{"EUR", "Euro", "Belgium", 978},
	Curr{"CLP", "Chilean Peso", "Chile", 152},
	...
}

func find(name string) {
	for i := 0; i < 10; i++ {
		c := currencies[i]
		switch {
		case strings.Contains(c.Currency, name),
		strings.Contains(c.Name, name),
		strings.Contains(c.Country, name):
		fmt.Println("Found", c)
		}
	}
}

// type switches
// go offers the interface{} type (or empty) as a super type implemented by all other types
func findAny(val interface{}) {
	switch i := val.(type) {
	case int:
		findNumber(i)
	case string:
		find(i)
	default:
		fmt.Printf("Unable to search with type %T\n", val)
	}
}


// for statements
// initialization, conditional expression, update statement
for x := 0; x<10; x++ {

}

for i, val := range values {

}

// loop over array, slice, map, channel
func main() {
	vals := []int{2,4,6}
	for _, v:= range vals {
		v--
	}

	fmt.Println(vals)
}

// break, continue, goto statements

// The goto statement is more flexible, in that it allows flow control to be transferred to an
// arbitrary location, inside a function, where a target label is defined. The goto statement
// causes an abrupt transfer of control to the label referenced by the goto statement.





