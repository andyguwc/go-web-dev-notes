/* Logging to a network

Logging to a network offers compelling advantages:
- Logs from many services can be aggregated to one central location.
- In the cloud, servers with only ephemeral storage can still have logs preserved.
- Security and auditability are improved.
- You can tune log servers and app servers differently.

*/



package main

import (
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:1902")
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)

	//logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile)

	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.")
}