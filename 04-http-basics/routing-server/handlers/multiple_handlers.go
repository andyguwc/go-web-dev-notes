
// When a path is resolved, it tries to go from the most specific to the least specific. In this case, any path that isn’t resolved prior to the / path will resolve to this one.
// It’s worth noting that paths ending in / can have redirection issues. In this listing, a user who visits /goodbye will be automatically redirected to /goodbye/. If you have query strings, they may be dropped. For example, /goodbye?foo=bar will redirect to /goodbye/.

// The handler registered to /goodbye/ will be executed for /goodbye (with a redirect), /goodbye/, /goodbye/foo, /goodbye/foo/bar, and so on.
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye/", goodbye)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
	// get name from query string
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	// look in the path for the name 
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Goodbye ", name)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}

