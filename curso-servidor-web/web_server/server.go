package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (server *Server) Handle(path string, handler http.HandlerFunc) {
	server.router.rules[path] = handler
}

func (server *Server) AddMidleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}

func (server *Server) Listen() error {
	http.Handle("/", server.router)

	err := http.ListenAndServe(server.port, nil)
	if err != nil {
		return err
	}

	return nil
}
