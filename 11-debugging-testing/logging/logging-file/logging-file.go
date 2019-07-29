/* logging to a file
initialize a log.Logger and send log mesages to that 
*/

package main 

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.Create("./log.txt")
	defer logfile.Close()

	// passing information to log.Logger - first the io.writer, second the prefix for the log messages, third a list of flags determining format of log message
	logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile)
	logger.Println("This is a regular message")
	logger.Fatalln("This is a fatal error")
}

