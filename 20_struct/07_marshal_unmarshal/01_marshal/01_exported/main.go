package main

import (
	"encoding/json"
	"fmt"
)

// Person is my struct
type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Printf("%T \n", &bs)
	fmt.Println(string(bs))
}
