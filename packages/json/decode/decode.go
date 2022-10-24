package decode

import (
	"encoding/json"
	"fmt"
)

type response struct {
	PageNumber int      `json:"page"`
	Fruits     []string `json:"fruits"`
}

func Decode() {
	str := `{"page":1, "fruits": ["apple","peach"]}`
	res := response{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println("page: ", res.PageNumber, "fruits: ", res.Fruits)
}
