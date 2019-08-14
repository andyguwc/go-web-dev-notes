
// declaring variables 
// var <identifier list> <type>

var name, desc string
var radius int32
var mass float64
var active bool
var satellites []string

// initialized declaration can omit types
var name, desc = "Mars", "Planet"
var radius = 6755

var satellites = []string{
	"Phobos",
	"Deimos",
}

// or use the var declaration block for cleaner results

var (
	name string = "Earth"
	desc string = "Planet"
)

// short variable declaration
func main() {
	name := "Neptune"
	desc := "Planet"
}

// can only be used within a function block
// <identifier list> := <value list or initializer expressions>
// := cannot be used to update a previously declared variable. Updates to variables must be done with an equal sign

// variable scope - available on the block and sub blocks 

// When a variable is declared at package level (outside of a function or method block), it is globally visible to the entire package, not just to the source file where the variable is declared. This means a
// package-scoped variable identifier can only be declared once in a group of files that make up a package


// constants
// Constants are values that can be represented by a text literal in the language
// constants must be declared with value literal 
const <identifier list> type = <value list or initializer expressions>
const a1, a2 string = "Mastering", "Go"