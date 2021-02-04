package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkStatus(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkStatus(link, c)
		}(l)
	}
}

func checkStatus(url string, c chan string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR: unable to make request to " + url)
	} else {
		fmt.Println("SUCCESS: Received \"" + response.Status + "\" from " + url)
	}

	c <- url
}
