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

func WeekDayz() {
	days := make(map[int]string)
	fmt.Println()

	days[0] = "Monday"
	days[1] = "Tuesday"
	days[2] = "Wednesday"
	days[3] = "Thursday"
	days[4] = "Friday"
	days[5] = "Saturday"
	days[6] = "Sunday"

	fmt.Println(days)

	days[6] = "SUNDAY"
	days[7] = "Juernesday"
	fmt.Println(days)

	delete(days, 7)
	fmt.Println(days)
}

func MapzSlices() {
	students := make(map[string][]int)
	students["John"] = []int{1, 2, 3}
	students["Bob"] = []int{4, 5, 6}

	fmt.Println(students)
	fmt.Println(students["John"])
	fmt.Println(students["Bob"][1])

}
