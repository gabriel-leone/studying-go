package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Create a slice of websites to check
	websites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://microsoft.com",
		"http://bing.com",
		"http://yahoo.com",
	}

	c := make(chan string)

	// Iterate over the websites and check their status
	for _, website := range websites {
		// Check the status of the website
		go checkWebsite(website, c)
	}

	// Receive the results from the channel
	for w := range c {
		go func(website string) {
			time.Sleep(5 * time.Second)
			checkWebsite(website, c)
		}(w)
	}

}

func checkWebsite(website string, c chan string) {
	// Make an HTTP request to the website
	_, err := http.Get(website)
	if err != nil {
		fmt.Println(website, "might be down!")
		c <- website
		return
	}

	fmt.Println(website, "is up!")
	c <- website
}
