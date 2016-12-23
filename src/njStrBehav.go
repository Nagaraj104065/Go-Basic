package main

import (
	"fmt"
)

type rectangle struct {
	height int
	width  int
}

//PROCEDURAL WAY IMPLEMENTATION
func cal(r *rectangle, mult int) int {
	r.height *= mult
	r.width *= mult

}

//OOPS WAY IMPLEMENTATION
func (rect *rectangle) calc() int {
	return 2 * (rect.height * rect.width)
}
func main() {
	rec := new(rectangle)
	rec.height = 10
	rec.width = 10
	var r rectangle
	r.height = 1
	r.width = 2
	//solution for OOPS WAY IMPLEMENTATION
	fmt.Println("area of rectangle value 10 and 10", rec.calc())
	//solution for PROCEDURAL WAY IMPLEMENTATION
	cal(&r, 2)
}
/*

func RetriveAll(tName string) []Student {
	session, err := mgo.Dial("127.0.0.1:27017")
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		panic("Error in connecting databases")
	} else {
		fmt.Println("connected")
	}
	dbconn := session.DB("school").C("student1")
	var student []Student
	err = dbconn.Find(bson.M{}).All(&student)
	return student
}

//-------------------------retriving one-------------------//
func RetriveOne(tName string) Student {
	session, err := mgo.Dial("127.0.0.1:27017")
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		panic("Error in connecting databases")
	} else {
		fmt.Println("connected")
	}
	dbconn := session.DB("school").C("student1")
	student := Student{}
	err = dbconn.Find(bson.M{}).All(&student)
	return student
}
