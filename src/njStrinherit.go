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

func (obj1 *cseDept) depDetail() {
	fmt.Println("student  id is :", obj1.stuID)
}
func depDetail1(obj2 *cseDept) {
	obj2.depDetail()
}
func main() {
	ant := []string{"n", "a", "g", "a", "r", "a", "j", "N", "a", "G", "a", "R", "a", "J"}
	cse := new(cseDept)
	college := new(colg)
	cse.stuID = 0
	cse.profID = 0
	fmt.Println(cse, ant, college.stuID)
	cse.depDetail()
	college.depDetail()
	depDetail1(cse)

	//	depDetail1(college)
	/*overeidding
		depDetail1(college)


		What ever happened to Polymorphism?
	Inheritance and Polymorphism must be understood
	separately
	People mix data generalisation and behaviour generalisation
	and get into deep trouble
	Go avoids this cleanly

	*/
	for i := 0; ; i++ {
		if i > 13 {
			i = 0

			defer func() {
				str := recover()
				fmt.Println(str)
			}()
			panic(i)
		}
		fmt.Printf("%s\r", ant[i])
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("hi")

}
