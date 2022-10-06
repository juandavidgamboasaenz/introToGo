package cycles

import "fmt"

func MyCycles() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 0

	// for _, v := range v {

	// }
	for {
		fmt.Println(j)

		if j == 2 {
			j++
			continue
		}

		j++

		if j > 10 {
			break
		}
	}
}
