package viewsportions

import "fmt"

func ViewsPortions() {
	base := []int{1, 0, 3, 4, 5}
	fmt.Println("Base : ", base)

	view := base[1:3] //Includes base elements from 1 to 2 indexes
	fmt.Println("vista 1 :", view)

	view2 := base[2:] //from index 2 to the last
	fmt.Println("vista 2 :", view2)

	view3 := base[:3] //From first index to index 2
	fmt.Println("vista 3 :", view3)

	view4 := base[:] // From first to last
	fmt.Println("vista 4 :", view4)

}
