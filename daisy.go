// Package daisy allows for expressjs-like daisy chaining of middleware
package daisy

import "net/http"

type Middleware func(http.Handler) http.Handler

type Linker struct {
	mws []Middleware
}

func Chain(mws ...Middleware) *Linker {
	c := make([]Middleware, len(mws))
	copy(c, mws)
	return &Linker{mws: c}
}

func (l *Linker) To(h http.Handler) http.Handler {
	for i := len(l.mws) - 1; i >= 0; i-- {
		h = l.mws[i](h)
	}
	return h
}

func (l *Linker) ToFunc(fn http.HandlerFunc) http.Handler {
	return l.To(fn)
}
