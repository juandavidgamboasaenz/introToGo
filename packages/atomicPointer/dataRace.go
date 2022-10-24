package atomicpointer

import (
	"fmt"
	"time"
)

func ShowDataRaceConnection(p *ServerConn) {
	for {
		time.Sleep(10 * time.Second)
		fmt.Println(p, *p)
	}
}

func DataRace() {

	c := make(chan bool)
	p := ServerConn{ID: "first_conn"}

	go ShowDataRaceConnection(&p)

	go func() {
		for {
			time.Sleep(13 * time.Second)
			newConn := ServerConn{ID: "new_conn"}
			p = newConn
		}
	}()
	<-c
}
