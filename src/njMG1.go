package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Student1 struct {
	Id   int    `bson:"id"`
	Name string `bson:"name"`
}

func main() {
	session, _ := mgo.Dial("127.0.0.1:27017")
	c := session.DB("test").C("student1")
	c.Insert(&Student1{Id: 1, Name: "sathya"})
	result := Student1{}
	c.Find(bson.M{}).All(&result)
	fmt.Println(result)
}
