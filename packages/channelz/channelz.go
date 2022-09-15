package channelz

import (
	"fmt"
)

func MyChannels() {
	channel := make(chan string)

	go func() {
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}()

	msg := <-channel
	fmt.Print(msg)

}
