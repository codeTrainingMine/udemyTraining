package main

import (
	"fmt"
	"sort"
)

func main() {
	studyGroup := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(sort.StringSlice(studyGroup))
	fmt.Println(studyGroup)
}
