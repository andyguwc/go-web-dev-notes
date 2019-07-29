// take advantage of a channel 

package main 

import (
	"fmt"
	"time"
)

func printCount(c chan int) { // int type channel passed in 
	num := 0
	for num >= 0 {
		num = <-c // wait for value to come in
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int) // create a channel
	a :=[]int{4,5,6,7,0,-1}
	go printCount(c) // starts goroutine
	for _, v := range a {
		c <-v 
	}
	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of main")
}