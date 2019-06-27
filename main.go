package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		// add as many links as needed
		"http://coris.network",
		"http://coris-docs.cyphercore.io/",
		"http://regen.coris.network/",
		"http://talk.edduus.dev/",
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

func checkStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down.")
		c <- link
		return
	}

	fmt.Println(link, "is up.")
	c <- link
}
