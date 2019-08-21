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



# Bridge Pattern
It decouples an abstraction from its implementation so that the two can vary
independently. This cryptic explanation just means that you could even decouple the most basic form of functionality: decouple an object from what it does.

This way, we can change what an object does as much as we want.


# Proxy Pattern
The Proxy pattern usually wraps an object to hide some of its characteristics. These
characteristics could be the fact that it is a remote object (remote proxy), a very heavy object such as a very big image or the dump of a terabyte database (virtual proxy), or a restricted access object (protection proxy).


# Decorator Pattern
The Decorator design pattern allows you to decorate an already existing type with more functional features without actually touching it. 

The Decorator type implements the same interface of the type it decorates, and stores an instance of that type in its members. This way, you can stack as many decorators (dolls) as you want by simply storing the old decorator in a field of the new one.

Use cases:
- When you need to add functionality to some code that you don't have access to, or you don't want to modify to avoid a negative effect on the code, and follow the open/close principle (like legacy code)
- When you want the functionality of an object to be created or altered dynamically, and the number of features is unknown and could grow fast

Example decorator for HTTP handler (adding logging capabilities), aka server middleware

```
type BasicAuthMiddleware struct {
    Handler http.Handler
    User string
    Password string
}
```

The BasicAuthMiddleware middleware stores three fields--a handler to decorate like in the previous middlewares, a user, and a password, which will be the only authorization to access the contents on the server. The implementation of the decorating method will proceed as follows:

```
func (s *BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r
*http.Request) {
    user, pass, ok := r.BasicAuth()
    if ok {
        if user == s.User && pass == s.Password {
            s.Handler.ServeHTTP(w, r)
        }
        else {
            fmt.Fprintf(w, "User or password incorrect\n")
        }
    }
    else {
        fmt.Fprintln(w, "Error trying to retrieve data from Basic auth")
    }
}
```

In the Decorator pattern, we decorate a type dynamically. This means that
the decoration may or may not be there, or it may be composed of one or many types. If you remember, the Proxy pattern wraps a type in a similar fashion, but it does so at compile time and it's more like a way to access some type.


In this aspect, you may think that the Proxy pattern is less flexible, and it is. But the Decorator pattern is weaker, as you could have errors at runtime, which you can avoid at compile time by using the Proxy pattern. Just keep in mind that the Decorator is commonly used when you want to add functionality to an object at runtime, like in our web server. It's
a compromise between what you need and what you want to sacrifice to achieve it.

# Facade Pattern
You use Facade when you want to hide the complexity of some tasks, especially when most of them share utilities (such as authentication in an API). A library is a form of facade, where someone has to provide some methods for a developer to do certain things in a friendly way. This way, if a developer needs to use your library, he doesn't need to know all the inner tasks to retrieve the result he/she wants.

Use the Facade design pattern in the following scenarios:
- When you want to decrease the complexity of some parts of our code. You hide
that complexity behind the facade by providing a more easy-to-use method.
- When you want to group actions that are cross-related in a single place.
- When you want to build a library so that others can use your products without
worrying about how it all works


# Flyweight Pattern
Thanks to the Flyweight pattern, we can share all possible states of objects in a single
common object, and thus minimize object creation by using pointers to already created
objects.