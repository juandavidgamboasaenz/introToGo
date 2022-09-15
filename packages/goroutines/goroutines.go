package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MyGoroutines() {
	// In this example is showed that a go routine executes code despite the order and
	// the other routines will be executed while the first is being executed aswell

	go mySlowName("David")

	fmt.Println("Que aburrido xdxd")
	var wait string
	fmt.Scanln(&wait)
	mySlowName(wait)
}

func mySlowName(name string) {
	fmt.Println("stuf")
	letters := strings.Split(name, "")
	fmt.Println("stuff")
	for _, letter := range letters {
		//If we put go before the timesleep all the time outs will be ignored,
		// but logical secuence will be still functioning
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
}
