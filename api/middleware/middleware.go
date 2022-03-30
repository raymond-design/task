package middleware

import (
	"net/http"
)

// creates a new instance of Middleware chains
func New(ms ...func(h http.Handler) http.Handler) *Middleware {
	return &Middleware{
		functions: ms,
	}
}

// HTTP router middleware type
type Middleware struct {
	functions []func(h http.Handler) http.Handler
}

// runs the request through the middleware chain, then serves it
func (m *Middleware) Then(h http.Handler) http.Handler {
	if h == nil {
		h = http.DefaultServeMux
	}
	for i := range m.functions {
		h = m.functions[len(m.functions)-1-i](h)
	}
	return h
}
