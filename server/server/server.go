package main

import "net/http"

// Server represents HTTP server
type Server struct {
	http.Server
}

// NewServer returns a new http server
func NewServer(addr string, configs ...func(*Server)) (*Server, error) {
	s := &Server{http.Server{Addr: addr}}
	for _, config := range configs {
		config(s)
	}
	return s, nil
}
