/* 
Implement TCP server from scratch 
https://github.com/GoesToEleven/golang-web-dev/tree/master/015_understanding-TCP-servers

*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	// returns a listener 
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	// accepts gives you a connections 
	// connection has a reader and a writer 
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// launch a go routine 
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	// read the message using scanner 
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()

	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")
}

