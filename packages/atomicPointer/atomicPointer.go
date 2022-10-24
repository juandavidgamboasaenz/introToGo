package atomicpointer

import (
	"fmt"
	"net"
	"sync/atomic"
)

type ServerConn struct {
	Connection net.Conn
	ID         string
	Open       bool
}

func AtomicPointer() {
	p := atomic.Pointer[ServerConn]{}
	s := ServerConn{ID: "first_conn"}
	p.Store(&s)

	fmt.Println(p.Load())

}
