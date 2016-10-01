package main

import (
	"errors"
	"fmt"
	"log"
)

// ErrNorgateMath ...
var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	fmt.Printf("type %T\n", ErrNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	// implementation
	return 42, nil
}
