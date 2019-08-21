/*
The bufio package offers several functions to do buffered writing of IO streams using an
io.Writer interface. The following snippet creates a text file and writes to it using buffered
IO


In general, the constructor functions in the bufio package create a buffered writer by
wrapping an existing io.Writer as its underlying source. For instance, the previous code
creates a buffered writer using the bufio.NewWriter function by wrapping the io.File
variable, fout

*/



package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rows := []string{
		"The quick brown fox",
		"jumps over the lazy dog",
	}

	fout, err := os.Create("./filewrite.data")
	writer := bufio.NewWriter(fout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fout.Close()

	for _, row := range rows {
		writer.WriteString(row)
	}
	writer.Flush()
}
