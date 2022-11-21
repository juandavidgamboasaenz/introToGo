package operatorz

import "fmt"

func Operate() {
	a := 4
	b := 5

	r := a == b

	fmt.Printf("%d is equal than %d? %t\n", a, b, r)

	r = a != b
	fmt.Printf("%d is not equal than %d? %t\n", a, b, r)

	r = a < b
	fmt.Printf("%d is less than %d? %t\n", a, b, r)

	r = a > b
	fmt.Printf("%d is greater than %d? %t\n", a, b, r)

	r = a <= b
	fmt.Printf("%d is less or equal than? %d %t\n", a, b, r)

	r = a >= b
	fmt.Printf("%d is greater or equal than? %d %t\n", a, b, r)

}
