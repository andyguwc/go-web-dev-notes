// exmaple to print concurrently 

package main 

import (
	"fmt"
	"time"
)

func count() {
	for i:=0; i<=5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond*1)
	}
}

func main() {
	// both count and main execute concurrently
	go count()
	time.Sleep(time.Millisecond*2)
	fmt.Println("Hello World")
	time.Sleep(time.Millisecond*2)
}