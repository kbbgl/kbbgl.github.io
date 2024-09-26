package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://google.com",
		"https://amazon.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}

	// for loop used with channels
	// will start loop when channel receives a value
	for link := range channel {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, channel)
		}(link)
	}

}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link) // blocking call

	if err != nil {
		fmt.Println(link, "might be down")
		channel <- link
		return
	}

	fmt.Println(link, "is up")
	channel <- link
}
