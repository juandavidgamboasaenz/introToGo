package slice

import "fmt"

type food int
type language int

//Here we use iota as a 0 type for language type int types
const (
	en language = iota
	fr
	es
	de
)

const (
	ap food = iota
	ta
	st
	bl
	pa
)

//If we need to access specific positions in an array slicing with :
//Is an easy way to do it
func Slicing() {
	arr := [3]int{1, 2, 3}

	sl02 := arr[0:2]
	sl12 := arr[1:2]
	sl23 := arr[2:3]
	fmt.Println(sl12)
	fmt.Println(sl23)
	fmt.Println(sl02)

}

func SliceCopy() {
	ar1 := []int{1, 2, 3}
	ar2 := [3]int{}
	copy(ar2[0:3], ar1)
	fmt.Println(ar2)

	slice := []int{1, 2, 3, 4, 5, 6}
	//We can create a slice, map or chan
	copia := make([]int, len(slice), cap(slice)*2)
	copy(copia, slice)

	fmt.Println(slice)
	fmt.Println(copia)
}

func indexedSlices() {

	option := [...]string{
		en: "English",
		fr: "French",
		es: "Spanish",
		de: "German",
	}

	fmt.Println(es, option[es])

	menu := [...]string{
		ap: "Apple",
		ta: "Tangerine",
		st: "Strawberry",
		bl: "BlueBerry",
		pa: "Papaya",
	}

	//When we are declaring a new slice, using a index to alocate the
	//data makes everything easier, at least if you want to index a data
	//thats not in the propper slice. Also you can declare specific data
	//positions into the array

	for _, v := range menu {
		fmt.Println("Todays menu is", v)
	}

	a4 := [...]int{10: 69420}
	fmt.Println(a4)

}

func MySlices() {
	//When declarin an array of fixed length
	var ages [3]int
	fmt.Println(ages)
	for _, v := range ages {
		fmt.Println(v)
	}
	fmt.Println(len(ages))

	var a1 [3]int = [3]int{10, 20, 30}
	fmt.Println(a1)

	var languages = [3]string{"go", "php", "javascript"}
	fmt.Println(languages)

	var a2 = [...]int{10, 20, 30}
	fmt.Println(a2)

	indexedSlices()

}

func VecPorc(vec [3]int, porc []int) {

	/*
		In this implementation using a vector and a portion
		shows the difference between a value re asign and a
		reference re assign, if the mokings are declared out
		side the functions the vector "v" will not save the
		reassigned value, besides the "p" value will re refe
		rence to the original portion without pointers notation0
			//Mockings
				v := [3]int{1, 2, 3}
				p := []int{1, 2, 3}
	*/

	vec[0] = 0
	if len(porc) > 0 {
		porc[0] = 0
	}

}
