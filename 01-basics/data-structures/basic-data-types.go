// Numeric Types

// To avoid any sort of confusion when porting Go on different platforms, the name of a numeric type reflects its
// size requirement. For instance, type int16 indicates an integer type that uses 16 bits for
// internal storage. 

// Types for storing integral values.
byte, int, int8, int16, int32,
int64, rune, 
// unsighed integers 
uint, uint8, uint16,
uint32, uint64, uintptr

// Types for storing floating point decimal values.
float32, float64 

// 

// Array [n]T An ordered collection of fixed size n of numerically indexed sequence of
// elements of a type T.

// Slice []T A collection of unspecified size of numerically indexed sequence of elements
// of type T.

// struct{} A structure is a composite type composed of elements known as fields (think
// of an object).

// map[K]T An unordered sequence of elements of type T indexed by a key of arbitrary
// type K.

// interface{} A named set of function declarations that define a set of operations that can be
// implemented by other types.


// Boolean type

bool

// String
// impmented as a slice of immutable byte values 


// Type Declaration
// bind a type to an indentifier to create a new named type that can be referenced

type <name identifier> <underlying type name>

type gallon float64


// Type conversion 

var test int32 = int32(actual) + count

