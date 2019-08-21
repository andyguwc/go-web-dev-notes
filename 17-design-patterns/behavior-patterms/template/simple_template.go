/* Template Pattern
Let user write a part of the program while the rest is executed by the abstraction

Example - three steps where the second step is deferred to the user

1. Each step in the algorithm must return a string.
2. The first step is a method called first() and returns the string hello.
3. The third step is a method called third() and returns the string template.
4. The second step is whatever string the user wants to return but it's defined by the
MessageRetriever interface that has a Message() string method.
5. The algorithm is executed sequentially by a method called ExecuteAlgorithm
and returns the strings returned by each step joined in a single string by a space.

*/


package template

import "strings"

type MessageRetriever interface {
	Message() string
}

type Templater interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

//-------------------------------------------------------------------------

// template implementor accept a MessageRetrisveer interface to execute as part of execution algorithm 
type Template struct{}

func (t *Template) first() string {
	return "hello"
}

func (t *Template) third() string {
	return "template"
}

// we define an ExecuteAlgorithm method that accepts the MessageRetriever interface as argument and returns the full algorithm: a
// single string done by joining the strings returned by the first(), Message() string and
// third() methods:

func (t *Template) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}

//-------------------------------------------------------------------------

type AnonymousTemplate struct{}

func (a *AnonymousTemplate) first() string {
	return "hello"
}

func (a *AnonymousTemplate) third() string {
	return "template"
}

// accept the func() method string type that we can implement directly 
func (a *AnonymousTemplate) ExecuteAlgorithm(f func() string) string {
	return strings.Join([]string{a.first(), f(), a.third()}, " ")
}

//---------------------------------------------------------------------------------

// using an adapter to maintain two templates one using MessageRetriever interface the other using the func 
// define adapter as private, as it shouldn't be used without a function defined in the myfunc field 
type adapter struct {
	myFunc func() string
}

func (a *adapter) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}

	return ""
}

func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &adapter{myFunc: f}
}


