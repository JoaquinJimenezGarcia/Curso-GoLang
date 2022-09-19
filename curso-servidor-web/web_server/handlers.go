package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "hello world!!!")
}
