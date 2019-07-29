/*

Accept arguments from command line to pass into a new file 
https://github.com/GoesToEleven/golang-web-dev/tree/master/003_string-to-html
*/


package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
		</head>
		<body>
		<h1>` +
		name +
		`</h1>
		</body>
		</html>
	`)
	// creaet new file 
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	// a new reader reading from s and then copy into the file 
	io.Copy(nf, strings.NewReader(str))
}