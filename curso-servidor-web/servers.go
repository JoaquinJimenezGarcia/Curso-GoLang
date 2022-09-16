package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	starting := time.Now()
	servers := []string{
		"https://platzi.com",
		"https://google.com",
		"https://instagram.com",
	}

	for _, server := range servers {
		checkServer(server)
	}

	totalTime := time.Since(starting)
	fmt.Printf("Total time: %s\n", totalTime)
}

func checkServer(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "it's not available currently.")
	} else {
		fmt.Println(server, "it's working normally.")
	}
}
