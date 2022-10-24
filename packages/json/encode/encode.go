package encode

import (
	"encoding/json"
	"fmt"
)

func Encode() {
	mapA := map[string]int{"apple": 5, "lettuce": 7}
	mapB, err := json.Marshal(mapA)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Printf(string(mapB))

	fmt.Printf(string(mapB))

}
