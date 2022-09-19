package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "hello world!!!")
}

func HandleHome(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "This is the API endpoint")
}
