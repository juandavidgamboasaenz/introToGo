package conditional

import "fmt"

func MyConditional() {
	b := true
	if b {
		fmt.Printf("condition is true so %t will be printed\n", b)
	} else {
		fmt.Printf("condition was declared false, so %t will be printed\n", b)
	}

	var n int
	fmt.Println("Ingrese un numero: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println(n, " es neutro")
	} else if n%2 == 0 {
		fmt.Println(n, " es par")
	} else if n%2 != 0 {
		fmt.Println(n, " es impar")

	}

}
