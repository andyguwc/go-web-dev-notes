# Server Synonymous Terms

https://github.com/GoesToEleven/golang-web-dev/tree/master/014_understanding-servers

router
request router
multiplexer
mux
servemux
server
http router
http request router
http multiplexer
http mux
http servemux
http server



# Handlers
A handler is an interface that has a method named ServeHTTP with two parameters: an HTTPResponseWriter interface and a pointer to a Request struct. In other words, anything that has a method called ServeHTTP with this method signature is a handler:
ServeHTTP(http.ResponseWriter, *http.Request)

ServeMux is an is an instance of Handler struct and DefaultServeMux is an instance of ServeMux


## Handler functions
Handler functions have the same signature as the
ServeHTTP method; that is, they accept a ResponseWriter and a pointer to a Request.

package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}

}


Go has a function type named HandlerFunc, which will adapt a function f with the appropriate signature into a Handler with a method f. For example, take the hello function:
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

If we do this:
helloHandler := HandlerFunc(hello)
then helloHandler becomes a Handler.


# Chaining Handlers


# Multiplexers
If handler is nil, default DefaultServeMux is used

Multiplexers are also handlers. ServeMux is an HTTP request multiplexer. It
accepts an HTTP request and redirects it to the correct handler according to
the URL in the request. DefaultServeMux is a publicly available instance of
ServeMux that is used as the default multiplexer.

- simple web server 

```
package main
import (
	"net/http"
)
func main() {
	http.ListenAndServe("", nil)
}
```

- specify server struct
```
package main
import (
	"net/http"
)
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServe()
}
```

the server struct config
type Server struct {
	Addr string
	Handler Handler
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	MaxHeaderBytes int
	TLSConfig *tls.Config
	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
	ConnState func(net.Conn, ConnState)
	ErrorLog *log.Logger
}



# Servemux
https://github.com/GoesToEleven/golang-web-dev/tree/master/018_understanding-net-http-ServeMux


type ServeMux
	func NewServeMux() *ServeMux
	func (mux *ServeMux) Handle(pattern string, handler Handler)
	func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
	func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)


# HTTPS
HTTPS is nothing more than layering HTTP on top of SSL (actually, Transport
Security Layer [TLS]). To serve our simple web application through HTTPS, we’ll use
the ListenAndServeTLS function

The cert.pem file is the SSL certificate whereas key.pem is the private key for the server

package main
import (
	"net/http"
)
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServeTLS("cert.pem", "key.pem")
}





