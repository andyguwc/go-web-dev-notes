# Singleton Pattern

As the name implies, it will provide you with a single instance of an object, and guarantee that there are no duplicates.

Use cases: 
- When you want to use the same connection to a database to make every query
- When you open a Secure Shell (SSH) connection to a server to do a few tasks, and don't want to reopen the connection for each task


# Builder Pattern

At the same time, you could be using the same technique to create many types of objects.
For example, you'll use almost the same technique to build a car as you would build a bus,
except that they'll be of different sizes and number of seats, so why don't we reuse the
construction process? This is where the Builder pattern comes to the rescue.

A Builder design pattern tries to:
- Abstract complex creations so that object creation is separated from the object user
- Create an object step by step by filling its fields and creating the embedded objects
- Reuse the object creation algorithm between many objects


We will create a director variable, the ManufacturingDirector type, to use the build processes represented by the product
builder variables for a car and motorbike. The director is the one in charge of construction of the objects, but the builders are the ones that return the actual vehicle. So our builder declaration will look as follows:

// use director variable (the ManufacturingDirector type)

type BuildProcess interface {
    SetWheels() BuildProcess
    SetSeats() BuildProcess
    SetStructure() BuildProcess
    GetVehicle() VehicleProduct
}

// preceding interface defines the steps that are necessary to build a vehicle
// we'll need a GetVehicle method to retrieve the Vehicle instance from the builder

type ManufacturingDirector struct {}

func (f *ManufacturingDirector) Construct() {
    //Implementation goes here
}


# Factory Method

Its purpose is to abstract the user from the knowledge of the struct he needs to achieve for a specific purpose, such as retrieving some value, maybe from a web service or a database. The user only needs an interface that provides him
this value. By delegating this decision to a Factory, this Factory can provide an interface that fits the user needs.

When using the Factory method design pattern, we gain an extra layer of encapsulation so that our program can grow in a controlled environment. With the Factory method, we delegate the creation of families of objects to a different package or object to abstract us from the knowledge of the pool of possible objects we could use.


# Prototype Method

While with the Builder pattern, we are dealing with repetitive building algorithms and with the factories we are simplifying the creation of many types of objects; with the Prototype pattern, we will use an already created instance of some type to clone it and complete it with the particular needs of each context.

The aim of the Prototype pattern is to have an object or a set of objects that is already created at compilation time, but which you can clone as many times as you want at runtime.


