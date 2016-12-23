package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sathya() {
	fmt.Println("function test")
}
func addition(a int, b int) int {
	return a + b
}

func closures(a int, b int) func() int {
	return func() (i int) {
		i = 1
		return i
	}
}

//class notes
func makeMultipliers(n int) func() int {
	i := 0
	return func() (retVal int) {
		retVal = n * i
		i++
		return
	}
}

///object and return function
func randV() []int {
	//return array value
	vall := make([]int, 5)
	vall[0] = rand.Intn(10)
	vall[2] = rand.Intn(10)
	return vall

}

// operator to the function
func checks(sum func(int, int) int, a int, b int) int {
	return sum(a, b)
}
func init() {
	//set the seed with time
	rand.Seed(time.Now().UTC().UnixNano())
}
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
	/// map Practicing
	man := make(map[string]int)
	man["sathya"] = 1
	var mapV map[string]int
	mapV = make(map[string]int)
	mapV["nj"] = 2
	mapV["rock"] = 3
	fmt.Println("map declaration", man["sathya"], mapV["nj"])
	val, found := man["nagaraj"]
	fmt.Println("Map Acces Test", val, found)
	fmt.Println("mapV", mapV)
	//delete keyword
	delete(mapV, "rock")
	fmt.Println("after deleted rock key", mapV)
	if _, present := mapV["rock"]; present {
		fmt.Println("Rock Key is Present")
	} else {
		fmt.Println("Rock Key not Present")
	}
	maps := map[string]int{"sathya": 1,
		"nagaraj": 2,
		"shakthi": 3}
	fmt.Println("lenght of the maps variable is", len(maps), maps["nagaraj"])
	var indexx string
	fmt.Println("Enter the string to search in maps ")
	fmt.Scanf("%s", &indexx)
	for index, value := range maps {
		fmt.Println("index valuse", index)
		if _, locked := maps[indexx]; locked {
			fmt.Println("found Nagaraj", index)
			break
		} else {
			fmt.Println(index, "=>", value)
		}
	}
	//function call
	sathya()
	fmt.Println(addition(2, 4))
	//testing anonymious function C01o5uR3
	sathyaa := func() string {
		return "sathya"
	}
	fmt.Println("testing anonymious function", sathyaa())
	val1 := makeMultipliers(3)
	val2 := closures(1, 3)
	fmt.Println("testing anonymious function with multiple return values", val1(), val2())
	fmt.Println("random value", randV())
	summ := func(a int, b int) int {
		return a + b
	}
	fmt.Println("function passing as argument", checks(summ, 1, 2))
}
