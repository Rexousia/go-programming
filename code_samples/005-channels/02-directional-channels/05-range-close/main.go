package main

import "fmt"

func main() {
	c := make(chan int)

	// SEND
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	//RECEIVE
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
