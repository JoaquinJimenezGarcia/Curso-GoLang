package main

import (
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (router *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := router.rules[path]

	return handler, exist
}

func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler, exist := router.FindHandler(request.URL.Path)
	if !exist {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	handler(writer, request)

}
