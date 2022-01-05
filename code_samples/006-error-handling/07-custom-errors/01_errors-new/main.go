package main

import (
	"errors"
	"log"
)

func main() {
	_, err := sqroot(-10)
	if err != nil {
		log.Fatalln(err)
	}

}

func sqroot(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate Math: square root of negative number")
	}
	return 42, nil

}
