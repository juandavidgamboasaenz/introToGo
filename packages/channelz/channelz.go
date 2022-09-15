package channelz

import (
	"fmt"
)

func MyChannels() {

	// A channel is a conduit thorugh which you can
	//send and receive values with the channel operator, <-
	channel := make(chan string)

	go func(channel chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}(channel)

	msg := <-channel
	fmt.Println(msg)

	for {
		msg = <-channel
		fmt.Println(msg)
	}

}
