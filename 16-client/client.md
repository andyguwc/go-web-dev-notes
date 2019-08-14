# Command Line Tool

## Flag library

```
// flagExample.go
package main
import (
"flag"
"log"
)
var name = flag.String("name", "stranger", "your wonderful name")
func main(){
flag.Parse()
log.Printf("Hello %s, Welcome to the command line world", *name)
}

// go build flagExample.go
// ./flagExample -name=Adam
// ./flagExample -name Adam

```

Bind variables through the init() function

var name String
func init() {
    flag.IntVar(&name, "name", "stranger", "your wonderful name")
}


# grequests
The grequests package has methods for performing all REST actions. The preceding
program uses the Get function from the package. It takes two function arguments. The first
one is the URL of the API, and the second one is the request parameters object.

If the REST method is GET, the RequestOptions holds the Params property. If the method is a POST, the struct will have a Data property.
Whenever we make a request, we get a response back. Let us see the structure of the
response. From the official documentation, the response looks like this:

type Response struct {
    Ok bool
    Error error
    RawResponse *http.Response
    StatusCode int
    Header http.Header
}


The Ok property of response holds the information about whether a request is successful or
not. If something went wrong, an error will be filled into the Error property. RawResponse

There are a few functions in Response that are useful:
JSON
XML
String
Bytes

