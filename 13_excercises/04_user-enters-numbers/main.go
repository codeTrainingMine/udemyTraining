package main

import "fmt"

func main() {
	var smallNumber, largeNumber int
	fmt.Println("Enter small number:")
	fmt.Scan(&smallNumber)
	fmt.Println("Enter large number:")
	fmt.Scan(&largeNumber)
	fmt.Println("Remainder of large number divided by small number:",
		largeNumber%smallNumber)
}
