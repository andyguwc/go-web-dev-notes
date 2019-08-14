// go no concept of passing by value
// to achieve passing by reference

package main 
import "fmt"

func half(val *float64) {
	fmt.Printf("call half(%f)\n", *val)
	*val = *val / 2
}

func main() {
	num := 2.807770
	fmt.Printf("num=%f\n", num)
	half(&num)
	fmt.Printf("half(num)=%f\n", num)
}

