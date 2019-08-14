// composite literal spec https://golang.org/ref/spec#Composite_literals


// Go also supports composite types such as array, slice, and map. These types are designed to store indexed elements of values of a specified type.



// Array

// Arrays are static entities that cannot grow or shrink in size once they are declared with a
// specified length. Arrays are a great option when a program needs to allocate a block of
// sequential memory of a predefined size. When a variable of an array type is declared, it is
// ready to be used without any further allocation semantics.


var val [100]int
var days [7]string

var board = [4][2]int{
	{33, 23},
	{62, 2},
	{23, 4},
	{51, 88},
}

var histogram = [5]map[string]int {
	map[string]int{"A":12,"B":1, "D":15},
	map[string]int{"man":1344,"women":844, "children":577,...},
}

// Array traversal

func max(nums [size]int) int {
	res := nums[0]
	for _, val := range nums {
		if val > res {
			res = val
		}
	}

	return res 
}

// sometimes passing array variable 



// Slices 
// size of slice is not fixed

var metalloids = []metalloid{
	metalloid{"Boron", 5, 10.81},
	metalloid{"Silicon", 14, 28.085},
	metalloid{"Germanium", 32, 74.63},
	metalloid{"Arsenic", 33, 74.921},
	metalloid{"Antimony", 51, 121.760},
	metalloid{"Tellerium", 52, 127.60},
	metalloid{"Polonium", 84, 209.0},
}

_ metalloids 

var ids = []string{"string 1", "string 2"}


func main() {
	// struct composite
	planet := struct {
	name string
	diameter int
	}{"earth", 12742}
}


// slice expressions with capacity
// <slice_or_array_value>[<low_index>:<high_index>:max]
// slice can be intialized at running the built-in function make
func main() {
	months := make([]string, 6)
}

// length and capacity
// Given a slice, its length and maximum capacity can be queried, using the len and cap
func main() {
	var vector []float64
	fmt.Println(len(vector)) // prints 0, no panic
	h := make([]float64, 4, 10)
	fmt.Println(len(h), ",", cap(h))
}

// appending to slices
func main() {
	months := make([]string, 3, 3)
	months = append(months, "Jan", "Feb", "March",
	"Apr", "May", "June")
	months = append(months, []string{"Jul", "Aug", "Sep"}...)
	months = append(months, "Oct", "Nov", "Dec")
	fmt.Println(len(months), cap(months), months)
}

// copy slices. Copy the content of v slice into result 
// Both source and target slices must be the same size and of the same type
func clone(v []float64) (result []float64) {
	result = make([]float64, len(v), cap(v))
	copy(result, v)
	return
}


// strings as slices 
// Internally, the string type is implemented as a slice using a composite value that points to
// an underlying array of rune. This affords the string type the same idiomatic treatment given
// to slices.

// The slice expression on a string will return a new string value pointing to its underlying
// array of runes.

// The string values can be converted to a slice of byte (or slice of rune) as
// shown in the following function snippet, which sorts the characters of a given string:

func sort(str string) string {
	bytes := []byte(str)
	var temp byte
	for i := range bytes {
		for j := i + 1; j < len(bytes); j++ {
			if bytes[j] < bytes[i] {
			temp = bytes[i]
			bytes[i], bytes[j] = bytes[j], temp
			}
		}
	}
	return string(bytes)
}



// Map
// The Go map is a composite type that is used as containers for storing unordered elements of
// the same type indexed by an arbitrary key value.
// map[<key_type>]<element_type>

// Map initialization
// <map_type>{<comma-separated list of key:value pairs>}

var (
	histogram map[string]int = map[string]int{
		"Jan":100, "Feb":445, "Mar":514, "Apr":233,
		"May":321, "Jun":644, "Jul":113, "Aug":734,
		"Sep":553, "Oct":344, "Nov":831, "Dec":312,
	}
	table = map[string][]int {
		"Men":[]int{32, 55, 12, 55, 42, 53},
		"Women":[]int{44, 42, 23, 41, 65, 44},
	}
)

// Making Map
func main() {
	hist := make(map[int]string)
	hist["Jan"] = 100
	hist["Feb"] = 445
	hist["Mar"] = 514
}

// Go provides a way to explicitly test for the absence of an element by returning an optional Boolean value
// as part of the result of an index expression
func save(store map[string]int, key string, value int) {
	val, ok := store[key]
	if !ok {
		store[ke] = value
	} else {
		panic(fmt.Sprintf("Slot %d taken", val))	
	}
}

// Map functions
len(map)
delete(map, key)


// Maps as parameters
func main() {
	hist := make(map[string]int)
	hist["Jun"] = 644
	hist["Jul"] = 113
	remove(hit, "Jun")
	len(hist) // returns 1
}

func remove(store map[string]int, key string) error {
	_, ok := store[key]
	if !ok {
		return fmt.Errorf("Key not found")
	}
	delete(store, key)
	return nil
}


// struct type

var empty struct{}

var car struct {make, model string}

var node struct{
	edges []string
	weight int
}

// struct initialization
// struct variable can be explicitly initialized using a composite literal 
var currency = struct{
	name, country string
	code int
}{
	"USD", "U.S.",
	100,
}


// declaring named struct types

type person struct {
	name string
	address address
}

type adress struct {
	street string
	city, state string
	postal string 
}

func makePerson() person {
	addr := address{
		city: "G",
		state: "C",
		postal: "123"
	}
	return person{
		name: "v",
		address: addr, 
	}
}

// structs as parameters
// struct variables store actual values so a new copy of struct value is created
// whenever a struct variable is reassigned or passed in as a function parameter


type person struct {
	name string
	title string
}
func updateName(p person, name string) {
	p.name = name
}
func main() {
	p := person{}
	p.name = "uknown"
	updateName(p, "V")
}

// instead need to pass in a pointer
func updateName(p *person, name string) {
	p.name = name
}



