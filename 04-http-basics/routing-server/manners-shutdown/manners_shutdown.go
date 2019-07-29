// graceful shutdown 
// When a server shuts down, you’ll often want to stop receiving new requests, save any data to disk, and cleanly end connections with existing open connections.

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/braintree/manners"
)

func main() {
	// gets instance of a handler 
	handler := newHandler()

	// set up monitoring of operating system signals 
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutdown(ch)

	// starts web server 
	manners.ListenAndServe(":8080", handler)
}

func newHandler() *handler {
	return &handler{}
}

type handler struct{}

// handler responding to web requests
func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

// waits for shutdown signal and reacts 
func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}

