package convertor

import (
	"fmt"
	"strconv"
)

func VarConvertor() {
	age := 22

	//For formating an int
	ageStr := strconv.Itoa(age)
	fmt.Println("My string age is:", ageStr)

	ageInt, err := strconv.Atoi(ageStr)
	fmt.Println("My int age is", ageInt)

	if err != nil {
		fmt.Println(err)
	}

}
