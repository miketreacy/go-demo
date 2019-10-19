package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type logWriter struct{}

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

// simple http request util using manual bytestring
func fetch(url string, byteLen int) (int, string, string, error) {
	// if strings.ToUpper(methodStr) == "POST" {
	// 	res, err := http.Post(url, "application/json", body)
	// }
	fmt.Printf("\nSending request to %v", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nfetch: %v\n", err)
		return 500, "Failed to fetch", "", err
	}
	// make an empty byte slice of initial length 99999
	bs := make([]byte, byteLen)
	res.Body.Read(bs)
	bodyStr := string(bs)
	return res.StatusCode, res.Status, bodyStr, err
}

// simple http request util using io.Copy
func otherFetch(url string) (int64, error) {
	fmt.Printf("\nSending request to  %v", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nfetch: %v\n", err)
		os.Exit(1)
	}
	// custom implementation of Writer interface
	lw := logWriter{}

	// This function takes two args:
	//   - a value the implements the Writer interface (os.Stdout)
	//   - a value that implements the Reader interface
	return io.Copy(lw, res.Body)
}

// This func makes the logWriter type satisfy the Write interface
// (could just use os.Stdout!)
func (logWriter) Write(bs []byte) (int, error) {
	bl := 1000
	newbs := string(bs)[:bl]
	fmt.Println(newbs)
	fmt.Println("\nJust wrote this many bytes:", len(newbs))
	return len(newbs), nil
}

func timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
