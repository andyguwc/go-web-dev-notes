/*
The net.Listener interface uses the Accept method to block indefinitely until a new
connection arrives from a client. 

The following abbreviated code snippet shows a simple
server that returns the string "Nice to meet you!" to each client connection and then
disconnects immediately.
*/



package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte("Nice to meet you!"))
		conn.Close()
	}

}

// test via telnet
// go run listen0.go & telnet 127.0.0.1 4040

