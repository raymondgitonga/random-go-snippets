package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		fmt.Println("alaa")
		c <- 1 // Send a signal; value does not matter.
	}()
	doSomethingForAWhile()
	<-c // Wait for sort to finish; discard sent value.
}

func doSomethingForAWhile() {
	fmt.Println("doSomethingForAWhile")
	time.Sleep(30 * time.Millisecond)
}
