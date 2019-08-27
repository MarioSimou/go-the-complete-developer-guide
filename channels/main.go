package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.stackoverflow.com",
		"https://golang.org",
	}
	c := 0

	for _, link := range links {
		go checkLink(link, &c)
	}

}

func checkLink(link string, c *int) {
	res, err := http.Get(link)
	printError(err)

	if res.StatusCode >= 200 && res.StatusCode < 400 {
		*c++
	}
}

func printError(e error) {
	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(0)
	}
}
