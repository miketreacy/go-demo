package main

import (
	"fmt"
	"net/http"
	"time"
)

var links = []string{
	"http://google.com",
	"http://stackoverflow.com",
	"http://golang.org",
	"http://amazon.com",
	"http://twitter.com",
}

func serialWebChecker() int64 {
	// This checks the urls synchronously in series
	startTime := timestamp()
	fmt.Println()
	for _, url := range links {
		checkURLOnce(url, nil)
	}
	endTime := timestamp()
	timeMs := endTime - startTime
	return timeMs
}

func parallelWebChecker() int64 {
	// This checks the urls synchronously in series
	startTime := timestamp()
	fmt.Println()

	// create a new channel
	c := make(chan string)

	// create a go routine for each function call
	for _, url := range links {
		go checkURLOnce(url, c)
	}

	// wait for a message from each go routine you created
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	// log out the first value that we receive from the channel
	// fmt.Println(<-c)

	endTime := timestamp()
	timeMs := endTime - startTime
	return timeMs
}

func checkURLOnce(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		msg := fmt.Sprintf("%v might be down", url)
		if c != nil {
			c <- msg
			return
		}
		fmt.Println(msg)
	}
	msg := fmt.Sprintf("%v is up!", url)
	if c != nil {
		c <- msg
		return
	}
	fmt.Println(msg)
}

func pollURLs() {
	// create a new channel
	c := make(chan string)

	// create a go routine for each function call
	for _, url := range links {
		go pingURL(url, c)
	}

	// infinite loop
	// will ping the url again every time a go routine sends result through channel
	// for {
	// 	go pingURL(<-c, c)
	// }

	// more explicit syntax for an infinite loop over a range of go routines
	// will ping the url again every time a go routine sends result through channel
	// for url := range c {
	// 	go pingURL(url, c)
	// }

	for url := range c {
		// function literal
		go func(u string) {
			// make this go routine sleep for 5 seconds
			secs := 2
			fmt.Printf("\n...sleeping for %v seconds\n", secs)
			time.Sleep(2 * time.Second)
			pingURL(u, c)
			// NEVER close over vars in a go routine
			// ALWAYS explicitly pass outer variables into the go routine!
		}(url)
	}
}

func pingURL(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		msg := fmt.Sprintf("%v might be down", url)
		if c != nil {
			c <- url
		}
		fmt.Println(msg)
	}
	msg := fmt.Sprintf("%v is up!", url)
	if c != nil {
		c <- url
	}
	fmt.Println(msg)
}
