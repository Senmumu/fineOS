package syscall

import "syscall"

const (
	EPANIC syscall.Errno = 0xffffff
)

var handlers [512]Handler

func wakeup(lock *uintptr, n int)

type Handler func(req *Request)

type Request struct {
	tf   *trapFrame
	Lock uintptr
}

func (r *Request) NO() uintptr {
	return r.tf.NO()
}
