package main

import (
	"fmt"
	"io"
	"net/http"
)

type webWriter struct{}

func (webWriter) Write(response []byte) (int, error) {
	fmt.Println(string(response))
	return len(response), nil
}

func main() {
	response, err := http.Get("https://google.com")
	if err != nil {
		fmt.Printf("[Error]", err)
	}

	writer := webWriter{}
	io.Copy(writer, response.Body)

}
