package main

import "fmt"

func main() {
	b := (true && false) || (false && true) || !(false && false)
	fmt.Println(b) // should be true
}
