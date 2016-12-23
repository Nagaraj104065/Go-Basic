package main

import (
	"fmt"
	"time"
)

type cseDept struct {
	stuID  int
	profID int
}

type colg struct {
	name string
	dept cseDept
}

func (obj1 *cseDept) depDetail() {
	fmt.Println("student  id is :", obj1.stuID)
}
func main() {
	ant := [2]string{"|", "~"}
	cse := new(cseDept)
	college := new(colg)
	cse.stuID = 0
	cse.profID = 0
	fmt.Println(cse, ant, college.dept.stuID)
	cse.depDetail()
	college.dept.depDetail()

	for i := 0; ; i++ {
		if i > 1 {
			i = 0
		}
		fmt.Printf("%s\r", ant[i])
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("hi")
}
