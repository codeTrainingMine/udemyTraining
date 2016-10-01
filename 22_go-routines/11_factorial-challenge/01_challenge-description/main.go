package main

import "fmt"

func main() {
	f := factorial(50)
	fmt.Println("Total:", f)
}

func factorial(n int) uint64 {
	var total uint64 = 1
	for i := n; i > 0; i-- {
		total *= uint64(i)
	}
	return total
}
