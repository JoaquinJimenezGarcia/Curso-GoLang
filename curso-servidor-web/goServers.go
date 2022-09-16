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

	i := 0

	for {
		if i > 2 {
			break
		}

		for _, server := range servers {
			go checkServer(server, channel)
		}

		time.Sleep(1 * time.Second)
		fmt.Println(<-channel)

		i++
	}

	totalTime := time.Since(starting)
	fmt.Printf("Total time: %s\n", totalTime)
}

func checkServer(server string, channel chan string) {
	_, err := http.Get(server)
	if err != nil {
		channel <- server + " NOT available"
	} else {
		channel <- server + " available"
	}
}
