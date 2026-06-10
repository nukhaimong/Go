package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	url    string
	status string
	err    error
}

func checkWebsite(url string, ch chan Result) {
	res, err := http.Get(url)
	if err != nil {
		ch <- Result{url: url, status: "down", err: err}
		return
	}
	ch <- Result{url: url, status: "up", err: nil}
	defer res.Body.Close()

}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.wrong-test.com",
	}
	starts := time.Now()
	ch := make(chan Result)
	for _, url := range urls {
		go checkWebsite(url, ch)
	}

	for range urls {
		result := <-ch
		if result.err != nil {
			fmt.Println(result.url, result.status, "Error:", result.err)
			continue
		}
		fmt.Println(result)
	}

	fmt.Println("Time Taken:", time.Since(starts))
}
