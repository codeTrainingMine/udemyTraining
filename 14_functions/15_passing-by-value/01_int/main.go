package main

import "fmt"

func main() {
	age := 44
	changeMe(age)
	fmt.Println(age)
}

func changeMe(z int) {
	fmt.Println(z)
	z = 24
}
