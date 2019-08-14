/* 
Due to lack of inheritance, Go encourages use of composition 

*/

package composition

type Athlete struct{}

func (a *Athlete) Train() {
	println("Training")
}

func Swim() {
	println("Swimming!")
}

// The CompositeSwimmerA type has a MyAthlete field of type Athlete. It also stores a func() type.
type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    *func()
}

//--------------------------------------------------------------------------

type Trainer interface {
	Train()
}
type Swimmer interface {
	Swim()
}

type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

//--------------------------------------------------------------------------

type Animal struct{}

func (r *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

//--------------------------------------------------------------------------

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

//--------------------------------------------------------------------------

type Parent struct {
	SomeField int
}

type Son struct {
	P Parent
}

func GetParentField(p Parent) int {
	return p.SomeField
}
