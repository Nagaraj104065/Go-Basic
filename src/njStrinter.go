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
	cseDept
}

// for instead of overridden

type Intrface interface {
	depDetail()
}

func (obj1 *cseDept) depDetail() {
	fmt.Println("student  id is :", obj1.stuID)
}
func depDetail1(obj2 Intrface) {
	obj2.depDetail()
}
func main() {
	ant := []string{"|", "/", "\\", "/"}
	cse := new(cseDept)
	college := new(colg)
	cse.stuID = 0
	cse.profID = 0
	fmt.Println(cse, ant, college.stuID)
	cse.depDetail()
	college.depDetail()
	depDetail1(college)

	for i := 0; ; i++ {
		if i > 3 {
			i = 0

		}
		fmt.Printf("%s\r", ant[i])
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("hi")

}
