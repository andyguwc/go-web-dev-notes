/*
issue a panic 
*/

package main
import "errors"
func main() {
	panic(errors.New("Something bad happened."))
}