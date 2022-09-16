package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	starting := time.Now()
	channel := make(chan string)

	servers := []string{
		"https://platzi.com",
		"https://google.com",
		"https://instagram.com",
	}

	for _, server := range servers {
		go checkServer(server, channel)
	}

	for _, s := range servers {
		fmt.Println(s, <-channel)
	}

	totalTime := time.Since(starting)
	fmt.Printf("Total time: %s\n", totalTime)
}

func checkServer(server string, channel chan string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "it's not available currently.")
		channel <- server + " NOT available"
	} else {
		fmt.Println(server, "it's working normally.")
		channel <- server + " available"
	}
}
