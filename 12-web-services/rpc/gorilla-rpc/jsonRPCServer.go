/* 

JSON RPC Server Using Gorilla Package

Here having a JSON file on the server that ahs details of books. The client requests book information by making an HTTP request

When RPC server receives the request, it reads the file from the filesystem and parses it.

If the given ID matches any book, then the server sends the information back to the client in the JSON format.
*/

package main

import (
	jsonparse "encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Args struct {
	Id string
}

// using the gorilla RPC json package 
type Book struct {
	Id     string `"json:string,omitempty"`
	Name   string `"json:name,omitempty"`
	Author string `"json:author,omitempty"`
}

type JSONServer struct{}

func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book
	raw, readerr := ioutil.ReadFile("./books.json")
	if readerr != nil {
		log.Println("error:", readerr)
		os.Exit(1)
	}
	marshalerr := jsonparse.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error:", marshalerr)
		os.Exit(1)
	}
	// Iterate over JSON data to find the given book
	for _, book := range books {
		if book.Id == args.Id {
			*reply = book
			break
		}
	}
	return nil
}

func main() {
	// create new RPC server
	s := rpc.NewServer()
	// register the type of data requested as JSON
	s.RegisterCodec(json.NewCodec(), "application/json")
	// register server by creating a new JSON server
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}


// Client Post JSON with a book ID 
// curl -X POST \
// http://localhost:1234/rpc \
// -H 'cache-control: no-cache' \
// -H 'content-type: application/json' \
// -d '{
// 	"method": "JSONServer.GiveBookDetail",
// 	"params": [{
// 	"id": "1234"
// 	}],
// 	"id": "1"
// }'