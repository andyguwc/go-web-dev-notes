/*
Reading from standard input

the fmt.Scan, fmt.Scanf, and fmt.Scanln are used to read data from standard input file handle, os.Stdin

*/

package main

import "fmt"

func main() {
	var choice int
	fmt.Println("A square is what?")
	fmt.Print("Enter 1=quadrilateral 2=rectagonal:")

	n, err := fmt.Scanf("%d", &choice)
	if n != 1 || err != nil {
		fmt.Println("Follow directions!")
		return
	}
	if choice == 1 {
		fmt.Println("You are correct!")
	} else {
		fmt.Println("Wrong, Google it.")
	}
}


