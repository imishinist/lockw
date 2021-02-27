package lockw

import (
	"context"
	"io"
	"sync"
)

type Writer interface {
	io.Writer
	Lock(ctx context.Context) (unlock func())
}

var _ Writer = &LockWriter{}

type LockWriter struct {
	w io.Writer
	s sync.Mutex
}

func NewLockWriter(w io.Writer) Writer {
	return &LockWriter{
		w: w,
		s: sync.Mutex{},
	}
}

func (l *LockWriter) Write(p []byte) (n int, err error) {
	l.s.Lock()
	defer l.s.Unlock()
	return l.w.Write(p)
}

func (l *LockWriter) Lock(ctx context.Context) func() {
	l.s.Lock()
	return func() {
		l.s.Unlock()
	}
}
