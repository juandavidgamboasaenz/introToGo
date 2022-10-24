package mapz

import "fmt"

func Mapz() {

	m := make(map[string]int)

	m["Clearity"] = 2
	m["route"] = 66

	i := m["route"]
	j := m["root"]

	fmt.Println(i, j == 0)

}
