
// path handlers for complex path 

package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

func main() {

	// get an instance of a path-based router 
	pr := newPathResolver()

	// maps functions to paths 
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goodbye) // sets http server to use the router 
	http.ListenAndServe(":8080", pr)
}

// create initialized pathresolver 
func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

// adds path to internal lookup 
func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

// For pathResolver to work as a handler function, it needs to implement the ServeHTTP method and implicitly implement the HandlerFunc interface. The ServeHTTP method is where path resolving happens.
func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path // iterates over registered paths 

	for pattern, handlerFunc := range p.handlers {
		if ok, err := path.Match(pattern, check); ok && err == nil {
			// executes handler for a mapped path 
			handlerFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
		}
	}

	http.NotFound(res, req) // if no path matches, the page wasn't found 
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Goodbye ", name)
}
