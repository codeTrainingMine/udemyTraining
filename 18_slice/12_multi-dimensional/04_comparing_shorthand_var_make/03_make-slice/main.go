package main

import "fmt"

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	student[0] = "Todd"
	//students = append(students, student)
	students[0][0] = "Blam"
	//student = append(student, "Todd")
	//student[0] = "Todder"
	fmt.Println(student)
	fmt.Println(students)
}
