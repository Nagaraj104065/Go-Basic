package main

import (
	"fmt"
)

type student struct {
	id   int
	name string
}

func main() {
	var s student
	s.id = 1
	s.name = "sathya"
	fmt.Println("calling struct by s var", s.id, s.name)
	// Struct initiallistion
	stu := student{id: 2, name: "nagaraj"}
	fmt.Println("calling struct by s var", stu.id, stu.name)
	//STRUCT POINTER
	stud := new(student)
	stud.id = 3
	stud.name = "shakthi"
	fmt.Println("calling struct by s var", stud.id, stud.name)

}
