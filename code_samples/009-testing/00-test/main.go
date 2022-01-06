package main

import (
	"fmt"

	"github.com/Rexousia/go-programming/tree/master/code_samples/009-testing/00-test/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
