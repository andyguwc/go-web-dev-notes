# Composite Pattern

In the Composite design pattern, you will create hierarchies and trees of objects. Objects have different objects with their own fields and methods inside them. This approach is very powerful and solves many problems of inheritance and multiple inheritances.


Binary tree compositions

Store instances of itself in a field (because it's recursive we must use pointers)
type Tree struct {
    LeafValue int
    Right *Tree
    Left *Tree
}



Can also embed objects 




func GetParentField(p *Parent) int{
    fmt.Println(p.SomeField)
}


# Adapter Pattern

The Adapter pattern is very useful when, for example, an interface gets outdated and it's not possible to replace it easily or fast. Instead, you create a new interface to deal with the current needs of your application, which, under the hood, uses implementations of the old interface.

