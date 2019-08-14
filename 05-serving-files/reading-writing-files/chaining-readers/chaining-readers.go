/* 
Wrap an existing reader and use its stream as the source for the new implementation. 

alphaReader type is now a struct which embeds an io.Reader value

alphaReader filters the stream of data for alpha characters
*/

package main 

import (
	"fmt"
	"io"
	"os"
	"strings"
)
type alphaReader struct {
	src io.Reader
}

func NewAlphaReader(source io.Reader) *alphaReader {
	return &alphaReader{source}
}

// implements the io.Reader interface which takes an stream of bytes and return the number of bytes and error (if necessary)
func (a *alphaReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	count, err := a.src.Read(p) // p has the source data 
	if err != nil {
		return count, err
	}

	// filter out the non alphabetic characters 
	for i :=0; i <len(p); i++ {
		if (p[i] >= 'A' && p[i] <= 'Z') ||
			(p[i] >= 'a' && p[i] <= 'z') {
				continue
		} else {
			p[i] = 0 // nil value for byte 
		} 
	}

	return count, io.EOF
}

func main() {
	// file, _ := os.Open("./reader2.go")
	str := strings.NewReader("hello! what is the weather?")
	alpha := NewAlphaReader(str)
	io.Copy(os.Stdout, alpha)
	fmt.Println()
}
