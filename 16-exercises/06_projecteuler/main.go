package main

import "fmt"

func main() {
	var d bool
	for i := 1; i < 3000000000; i++ {
		for j := 1; j <= 20; j++ {
			if i%j > 0 {
				d = false
				break
			} else {
				d = true
			}
		}
		if d {
			fmt.Println("found:", i)
			break
		}

	}
}

// https://projecteuler.net/problem=5
// Smallest multiple
// Problem 5
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
