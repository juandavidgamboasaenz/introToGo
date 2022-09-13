package inout

import "fmt"

func MyInOut() {

	var (
		name string
		age  int
		pi   float64
	)

	fmt.Println("Input your name: ")
	fmt.Scanln(&name)

	fmt.Println("Input your age: ")
	fmt.Scanln(&age)

	fmt.Println("Input your pi value: ")
	fmt.Scanln(&pi)

	fmt.Printf("Name: %s Age: %d pi value %f\n", name, age, pi)
}
