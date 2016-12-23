package main

import "fmt"

type array []int
type number int

//RECIVER CREATE USER DEFINED TYPE

func (a array) val(fun func(int)) {
	for i := 0; i <= len(a)-1; i++ {
		fun(a[i])
	}
}
func (n number) soln() int {
	return int(n)
}
func main() {
	var arr array = []int{1, 2, 6, 4, 5}
	var num number = 1
	value := 0
	arr.val(func(va int) {
		value += va
	})
	fmt.Println(value, len(arr))
	fmt.Println(num.soln())
	fmt.Println("\x6e\x61\x67\x61\x72\x61\x6a")
	var name string
	fmt.Scanln(&name)
	fmt.Println(name)
}
