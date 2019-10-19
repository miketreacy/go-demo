package main

import (
	"fmt"
	"net/http"
)

var links = []string{
	"http://google.com",
	"http://stackoverflow.com",
	"http://golang.org",
	"http://amazon.com",
	"http://twitter.com",
}

func serialWebChecker() {
	// This checks the urls synchronously in series
	startTime := timestamp()
	fmt.Println()
	for _, url := range links {
		checkURL(url, nil)
	}
	endTime := timestamp()
	timeMs := endTime - startTime
	fmt.Println("\nserialWebChecker took", timeMs, "milliseconds to run")

}

func parallelWebChecker() {
	// This checks the urls synchronously in series
	startTime := timestamp()
	fmt.Println()

	// create a new channel
	c := make(chan string)
	for _, url := range links {
		go checkURL(url, c)
	}
	endTime := timestamp()
	timeMs := endTime - startTime
	fmt.Println("\nparallelWebChecker took", timeMs, "milliseconds to run")

}

func checkURL(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("\n", url, "might be down")
	}
	fmt.Println("\n", url, "is up")
}
