package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	//send
	go fooSend(c)

	//receive
	barReceive(c)

	fmt.Println("About to exit")
}

//send
func fooSend(c chan<- int) {
	c <- 42

}

//receive
func barReceive(c <-chan int) {
	fmt.Println(<-c)

}
