package main

import (
	"fmt"
	"time"
)

const strings string = "hI IaM NaGarAj"

func main() {
	i := 1122222
	j := 23222222
	s := "sathya"
	r := "nagaraj"
	var booli bool = true
	var gaali bool = false
	var k int
	k = i | j
	var l int = i & j
	var n, a, g, aa uint
	n, a, g, aa = 1, 2, 1, 2
	fmt.Println("Helllooooooooooooooooooooo", int64(i+j), s+r, gaali, booli, k, l, strings, n>>a, g<<aa)
	if n < a {
		fmt.Println("hid", time.Now())
	} else {
		fmt.Printf("hahahahaha")
	}
	//String
	var string string = "I aM sAthya"
	length := len(string)
	fmt.Println("String length is", length)
	for i, j := range string {
		fmt.Println("range examble", i, j)
	}
	//Patern for right arrow
	for ii := 1; ii <= 22; ii++ {
		if ii <= 10 {
			for jj := 1; jj <= ii; jj++ {
				fmt.Print("*")
				time.Sleep(100 * time.Millisecond)
			}
		} else if ii <= 20 {
			for jj := 20; jj >= ii; jj-- {
				fmt.Print("#")
				time.Sleep(100 * time.Millisecond)
			}
		} else {
			break
		}
		fmt.Println(ii)
	}
	//array
	arr := [5]int{1, 2, 3, 5, 7}
	fmt.Println(arr[2])
	//Patern for N LETTER
	for kk := 1; kk <= 7; kk++ {
		for ll := 1; ll <= 10; ll++ {
			kkk := kk + 1
			lll := ll
			if ll <= 3 || ll >= 8 && ll <= 10 || kk == ll || kkk == lll {
				fmt.Print("-")
				time.Sleep(100 * time.Millisecond)

			} else {
				fmt.Print(" ")
				time.Sleep(100 * time.Millisecond)
			}
		}
		fmt.Println("")
	}

}
