package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = "-12"
	var y = 6
	z, err := strconv.Atoi(x)
	fmt.Println(z)
	fmt.Println(y + z)
	fmt.Printf("%T \n", z)
	fmt.Println(err)

	//i := int(x) // gives error

}
