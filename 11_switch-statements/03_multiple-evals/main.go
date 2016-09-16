package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or, err Jenny")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "John":
		fmt.Println("Wassup John")
	default:
		fmt.Println("Have you no friends")
	}
}
