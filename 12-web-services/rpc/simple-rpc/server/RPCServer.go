
/* Create RPC Server

Arguments passed and value returned 

*/

package main

import (
	"log"
	"net"
	"net/rpc"
	"time"
	"net/http"
)

// first create Args struct which holds info about arguments passed from RPC to server
type Args struct{}

type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	// Fill reply pointer to send the data back
	*reply = time.Now().Unix()
	return nil
}

func main() {
	// create a new RPC server 
	timeserver := new(TimeServer)
	// register RPC server
	rpc.Register(timeserver)
	rpc.HandleHTTP()
	// start a TCP server to listen for requests
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}