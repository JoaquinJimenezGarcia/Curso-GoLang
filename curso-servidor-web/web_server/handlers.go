package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "hello world!!!")
}

func HandleHome(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "This is the API endpoint")
}

func UserPostRequest(writer http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User

	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(writer, "Error: %v")
		return
	}

	response, err := user.ToJson()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)
}
