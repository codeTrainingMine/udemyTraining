package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		fmt.Printf("%T\n", err)
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}
	// implementation
	return 42, nil
}
