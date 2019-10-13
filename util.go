package main

import (
	"fmt"
	"net/http"
	"os"
)

// Index returns the first index of the target string t, or -1 if no match is found.
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include returns true if the target string t is in the slice.

func includes(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any returns true if one of the strings in the slice satisfies the predicate f.
func any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if all of the strings in the slice satisfy the predicate f.
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter returns a new slice containing all strings in the slice that satisfy the predicate f.
func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map returns a new slice containing the results of applying the function f to each string in the original slice.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// simple http request util
func fetch(url string) (int, string, string) {
	// if strings.ToUpper(methodStr) == "POST" {
	// 	res, err := http.Post(url, "application/json", body)
	// }

	fmt.Printf("\nSending request to  %v", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	// make an empty byte slice of initial length 99999
	bs := make([]byte, 99999)
	res.Body.Read(bs)
	bodyStr := string(bs)
	// body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	return res.StatusCode, res.Status, bodyStr
}
