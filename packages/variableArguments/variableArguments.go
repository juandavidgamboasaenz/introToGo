package variablearguments

import "fmt"

func VariableArguments() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{1, 2, 3, 4}
	arr3 := []int{1, 2, 3, 4, 5}
	arr4 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Sum(arr1...))
	fmt.Println(Sum(arr2...))
	fmt.Println(Sum(arr3...))
	fmt.Println(Sum(arr4...))
}

func Sum(n ...int) int {
	/* Mockings!
	arr1 := []int{1, 2, 3}
	arr2 := []int{1, 2, 3, 4}
	arr3 := []int{1, 2, 3, 4, 5}
	arr4 := []int{1, 2, 3, 4, 5, 6}
	*/

	var total int
	for _, v := range n {
		total += v
	}
	return total
}
