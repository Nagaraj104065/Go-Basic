package main

import (
	"fmt"
)

type msg interface {
}

/* if there is declared by empty interface that may be any
data type or custom type */

type Body struct {
	Msg interface{}
}

func val(i msg) {
	fmt.Println(i)
}
func main() {
	b := Body{}
	b.Msg = "5"
	var i interface{}
	i = 0
	val(i)
	j := "sathya"
	val(j)
	fmt.Printf("%#v %T \n", b.Msg, b.Msg) // Output: "5" string
	b.Msg = 5
	fmt.Printf("%#v %T\n", b.Msg, b.Msg) //Output:  5 int
}
