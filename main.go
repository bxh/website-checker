package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// Create a Go routine.
		go checkLink(link, c)
	}

	// range can be used as for infinite stream too. This is equivalent to `for {...}`
	for l := range c {
		// Function Literal
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c) // if `link` is replaced with `l`, then `l` will be fixated at its initial value. Never share variable between goroutines directly.
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
