package errhandler

import (
	"fmt"
	"net/http"
)

func ErrHandling() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
