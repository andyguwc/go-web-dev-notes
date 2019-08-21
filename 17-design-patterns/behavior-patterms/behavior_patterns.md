# Strategy Pattern
The Strategy pattern uses different algorithms to achieve some specific functionality. These
algorithms are hidden behind an interface and, of course, they must be interchangeable. All
algorithms achieve the same functionality in a different way.

The objectives of the Strategy pattern are really clear. The pattern should do the following:
Provide a few algorithms to achieve some specific functionality
All types achieve the same functionality in a different way but the client of the
strategy isn't affected



Refator to create a library
Problem 1:
It cannot be used as a library. We have critical code written in the main package
(strategy creation). Solution: Abstract to two different packages the strategy
creation from the command-line application.

Solution:
First, we will place our main package and function
outside of the root package; in this case, in a folder called cli. It is also common to call this
folder cmd or even app. Then, we will place our PrintStrategy interface in the root
package, which now will be called the strategy package. Finally, we will create a shapes
package in a folder with the same name where we will put both text and image strategies.
So, our file structure will be like this:


Problem 2: 
None of the strategies are doing any logging to file or console. We must provide a
way to read some logs that an external user can integrate in their logging
strategies or formats. Solution: Inject an io.Writer interface as dependency to
act as a logging sink.

Solution:
We have added the SetLog(io.Writer) method to add a logger strategy to our types; this
is to provide feedback to users.


# Chain of Responsibility
As its name implies, it consists of a chain and, in our case, each link of the chain follows the single responsibility principle.

The single responsibility principle implies that a type, function, method, or any similar
abstraction must have one single responsibility only and it must do it quite well.

The objective of the chain of responsibility is to provide to the developer a way to chain
actions at runtime. The actions are chained to each other and each link will execute some
action and pass the request to the next link (or not). The following are the objectives
followed by this pattern:
- Dynamically chain the actions at runtime based on some input
- Pass a request through a chain of processors until one of them can process it, in which case the chain could be stopped


# Command Pattern
Connect types that are unrelated 

A Command pattern is commonly seen as a container. You put something like the info for
user interaction on a UI that could be click on login and pass it as a command. You
don't need to have the complexity related to the click on login action in the command
but simply the action itself

The command pattern will be used heavily when dealing with channels. With channels you
can send any message through it but, if we need a response from the receiving side of the
channel, a common approach is to create a command that has a second, response channel
attached where we are listening

When using the Command design pattern, we are trying to encapsulate some sort of action
or information in a light package that must be processed somewhere else.
- Put some information into a box. Just the receiver will open the box and know its
contents.
- Delegate some action somewhere else.


# Template Design Pattern

The Template design pattern is all about reusability and giving responsibilities to the user.
So the objectives for this pattern are following:
Defer a part of an algorithm of the library to a the user
Improve reusability by abstracting the parts of the code that are not common
between executions


# Momento Pattern

The meaning of memento is very similar to the functionality it provides in design patterns.
Basically, we'll have a type with some state and we want to be able to save milestones of its
state. Having a finite amount of states saved, we can recover them if necessary for a variety
of tasks-undo operations, historic, and so on.

Memento: A type that stores the type we want to save. Usually, we won't store
the business type directly and we provide an extra layer of abstraction through
this type.
Originator: A type that is in charge of creating mementos and storing the current
active state. We said that the Memento type wraps states of the business type and
we use originator as the creator of mementos.
Care Taker: A type that stores the list of mementos that can have the logic to
store them in a database or to not store more than a specified number of them.

# Interpreter Pattern
Solve business cases where it's useful to have a language to perform common operations

Designing a new language, big or small, can be a time consuming task so it's very important
to have the objectives clear before investing time and resources on writing an interpreter of
it:
Provide syntax for very common operations in some scope (such as playing
notes).
Have a intermediate language to translate actions between two systems. For
example, the apps that generate the Gcode needed to print with 3D printers.
Ease the use of some operations in an easier-to-use syntax.