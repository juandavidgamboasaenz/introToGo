package atomicpointer

import (
	"fmt"
	"sync/atomic"
	"time"
)

func AtomicPointerUseCase() {
	c := make(chan bool)
	p := atomic.Pointer[ServerConn]{}
	s := ServerConn{ID: "first_conn"}
	p.Store(&s)

	go ShowUseCaseConnection(&p)

	go func() {
		for {
			time.Sleep(13 * time.Second)
			newConn := ServerConn{ID: "new_conn"}
			p.Swap(&newConn)
		}
	}()

	<-c
}

func ShowUseCaseConnection(p *atomic.Pointer[ServerConn]) {
	for {
		time.Sleep(10 * time.Second)
		fmt.Println(p, p.Load())
	}
}
