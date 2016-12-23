package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	///Author
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
	fmt.Println("Programmed By NAGARAJ")
	//switch
	key := 2
	switch key {
	case 1:
		fmt.Println("Hi")
	case 2:
		fmt.Println("Hello")
		fallthrough
	case 3:
		fmt.Println("world")
		fallthrough
	case 4:
		fmt.Println("Nagaraj")
		break
	case 5:
		fmt.Println("sathya")
		fallthrough
	default:
		fmt.Println("NagA")
	}
	//String
	var string string = "I aM sAthya"
	length := len(string)
	fmt.Println("String length is", length)
	for i, j := range string {
		fmt.Println(j, strconv.Itoa(i))
	}
	fmt.Println("Hiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiihelooooooooooooooooooooo")
	fmt.Println("Hii\\n\tI\tam\t\"nagaraj\"\n")
	fmt.Println("string[1:]", string[1:], "string[1:7]", string[1:7], "string[:]", string[:], "string[1:7]", string[:8])
	//Pointer
	var varPoint *int
	jk := 3
	varPoint = &jk
	fmt.Println("Pointer var address =>variable address", varPoint, &jk)
	fmt.Println("Pointer var value =>variable value", *varPoint, jk)
	//Slice Practicing
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println("slice 1 valuee", slice)
	slice2 := make([]int, 3, 10)
	slice2 = slice
	slice2 = []int{1, 2, 3, 4, 4}
	fmt.Println("slice 2 value", slice2)
	ars := [5]int{1, 2, 3, 4, 5}
	var sliceA []int
	sliceA = ars[:]
	fmt.Println("Slice Vs Array", sliceA, ars)
	//Changing Slice value
	sliceA[0] = 3
	sliceA[1] = 3
	fmt.Println("after chenged the value Slice Vs Array", sliceA, ars)
	//Slice Append
	sliceA2 := append(sliceA)
	sliceA2[0] = 111
	fmt.Println("slice Append  value", sliceA2)
	fmt.Println("after Append the value Slice", sliceA2, sliceA, ars)
	//Changing SliceA2 value
	sliceA2[0] = 3
	sliceA2[1] = 3
	fmt.Println("after changed the value SliceA2", sliceA2, sliceA, ars)
	//Slice Copy
	var sliceC = make([]int, 5, 10)
	copy(sliceC, sliceA2)
	fmt.Println("Coping the SliceA2,SliceC", sliceA2, sliceC)
	/* Practicing Array */
	var array = [5]int{1, 2, 3, 4, 5}
	array2 := [5]int{13,
		45,
		6,
		7,
		7}
	var array3 [5]int
	var array4 [5]int
	fmt.Println(array[2])
	fmt.Println(len(array) > len(array2))
	for i := 0; i <= len(array)-1; i++ {
		array3[i] = array[i] + array2[i]
		array4[i] = array3[i] + i
		fmt.Println(array4[i])
	}
	fmt.Print(array3)
	// Two Dimentional array//
	arr := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Print(arr)
	var arr2 [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scanf("%d", &arr2[i][j])
		}
		fmt.Println("")
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(arr2[i][i])
		}
		fmt.Println("")
	}

}
