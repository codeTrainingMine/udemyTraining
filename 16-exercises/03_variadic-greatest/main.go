package main

import "fmt"

func main() {
	greatest(1, 2, 31212, 4, 5, 6)
}

func greatest(xn ...int) {
	var holder int
	for _, v := range xn {
		if v > holder {
			holder = v
		}
	}
	fmt.Println(holder)
}
