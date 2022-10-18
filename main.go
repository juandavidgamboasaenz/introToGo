package main

import (
	"fmt"

	"github.com/LordCeilan/introToGo/packages/slice"
)

func main() {

	v := [3]int{1, 2, 3}
	p := []int{1, 2, 3}
	slice.VecPorc(v, p)
	fmt.Println("vector: ", v)
	fmt.Println("portion: ", p)
}
