package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m["Todd"])
	changeMe(m)
	fmt.Println(m["Todd"])
}

func changeMe(z map[string]int) {
	z["Todd"] = 44
}
