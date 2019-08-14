// Go Methods
// When a function is scoped to a type, or attached to the type, it is known as a method. A method is
// defined just like any other Go function. However, its definition includes a method receiver,
// which is an extra parameter placed before the method's name, used to specify the host type
// to which the method is attached.

// As mentioned, a method has the scope of a type. Therefore, it can only be accessed via a
// declared value (concrete or pointer) of the attached type using dot notation.

type gallon float64

func (g gallon) quart() quart {
	return float64(g * 4)
}

func main(){
	gal := gallon(5)
	fmt.Println(gal.quart())
}


// objects in go
// Object: A data type that stores states and exposes behavior
// In Go all types can achieve this. There is no special type called a
// class or object to do this. Any type can receive a set of method to
// define its behavior, although the struct type comes the closest to
// what is commonly called an object in other languages.

// Composition
// Using a type such as a struct or an interface (discussed later), 
// it is possible to create objects and express their polymorphic relationships through composition.

// Type Inheritance
// Go does not support polymorphism through inheritance. A newly
// declared named type does not inherit all attributes of its
// underlying type and are treated differently by the type system. As
// a consequence, it is hard to implement inheritance via type lineage
// as found in other languages.

// The struct type, however, offers all of the features that are traditionally attributed to objects in other languages, such as:
// Ability to host methods
// Ability to be extended via composition
// Ability to be sub-typed (with help from the Go interface type)

// Go uses the composition over inheritance principle to achieve polymorphism using the type
// embedding mechanism supported by the struct type. In Go, there is no support for
// polymorphism via type inheritance

func main() {
	t := &truck{
		vehicle: vehicle{}
	}
}

// constructor function 
// converntional idiom is to use factory function to create and initialize value for a type
// While not required, providing a function to help with the initialization of composite values,
// such as a struct, increases the usability of the code.
type truck struct {
	vehicle
	wheels int
}

func newTruck(mk, mdl string) *truck {
	return &truck{vehicle:vehicle{mk, mdl}}
}


