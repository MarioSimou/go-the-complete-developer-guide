package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	fmt.Println(res.Status)
	fmt.Println(res.StatusCode)
	fmt.Println(res.Body)
}

func Read(p []byte) (n int, err error) {
	fmt.Println(p)
}
