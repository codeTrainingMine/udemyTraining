package main

import "fmt"

func main() {
	for i := 1; i <= 1800; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}
