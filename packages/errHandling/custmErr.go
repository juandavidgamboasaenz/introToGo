package errhandler

import (
	"errors"
	"fmt"
)

func Increment(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("math: cannot process negative number")
	}

	return (n + 1), nil
}

func CustmErr() {
	num := 5

	if inc, err := Increment(num); err != nil {
		fmt.Printf("Failed Number: %v", inc)
	} else {
		fmt.Printf("Incremented Number: %v", inc)
	}

}
