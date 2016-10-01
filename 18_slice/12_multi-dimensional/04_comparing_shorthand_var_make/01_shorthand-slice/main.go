package main

import "fmt"

func main() {
	student := []string{}
	students := [][]string{}
	//student[0] = "Todd"
	student = append(student, "Todd")
	student[0] = "Todder"
	fmt.Println(student)
	fmt.Println(students)
}
