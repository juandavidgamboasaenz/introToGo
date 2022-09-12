package types

import "fmt"

func MyTypes() {
	var b1, b2, b3 int
	var s string = "my string"
	var i int = 1
	var t bool = true
	sl := []int{1, 2, 3, 4, 5}
	fmt.Printf("String : %s, Integer : %d, Boolean :%t, Slice: %v\n", s, i, t, sl)
	fmt.Println(b1, b2, b3)

	var a2 int = 2
	var c2, c3 int = -15, 0
	var d1, d2, d4 bool = true, false, true
	fmt.Println(a2, c2, c3, d1, d2, d4)

	var e1, e2 = 1, 2
	var f1, f2 = true, false
	var h1 = []string{"nya", "nyanyame", "nya"}
	fmt.Println(e1, e2, f1, f2, h1)

	k1, k2 := 1, 2
	l1, l2 := true, false
	m1, m2 := []int{1, 2, 3}, []string{"s", "s2"}
	fmt.Println(k1, k2, l1, l2, m1, m2)

	var (
		o  = 1
		n1 = true
		j  = []bool{true, false}
	)
	fmt.Println(o, n1, j)

}
