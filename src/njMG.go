package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func main() {
	sessions, err := mgo.Dial("localhost:27017,localhost:27018,localhost:27019")
	if err != nil {
		panic(err)
	}
	c := sessions.DB("mydb").C("testCollection")
	c.Insert(&Person{Name: "Foo", Age: 120})
	var result []Person
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	c.Remove(bson.M{"name": "Foo"})
}
