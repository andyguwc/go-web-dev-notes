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


