/* 
Custom Middlewares (http wrapper)

- Create a handler function by passing the main handler function (mainLogic)
to http.HandlerFunc().

- Create a middleware function that accepts a handler and returns a handler.

- The method ServeHTTP allows a handler to execute the handler logic that
is mainLogic.

- The http.Handle function expects an HTTP handler. By taking that into
consideration, we wrapped up our logic in such a way that, finally, a handler gets
returned, but the execution is modified.

- We are passing the main handler into the middleware. Then middleware takes it
and returns a function while embedding this main handler logic in it. This makes
all the requests coming to the handler pass through the middleware logic.

*/

package main 

import (
	"fmt"
	"net/http"
)

// note that http.HandlerFunc is a valid http.Handler too 
func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase!")
		handler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after response phase!")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	// Business logic goes here
	fmt.Println("Executing mainHandler...")
	w.Write([]byte("OK"))
}

func main() {
	// HandlerFunc returns a HTTP Handler
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mainLogicHandler))
	http.ListenAndServe(":8000", nil)
}


