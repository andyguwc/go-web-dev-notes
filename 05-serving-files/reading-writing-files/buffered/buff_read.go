/*
Reading buffered streams can be done simply by calling the constructor function
bufio.NewReader to wrap an existing io.Reader
*/


package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./bufread0.go")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {

	// The previous code uses the reader.ReadString method to read a text file using the '\n' character as the content delimiter.
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Error reading:, err")
				return
			}
		}
		fmt.Print(line)
	}
}

