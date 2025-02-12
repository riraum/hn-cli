package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello hn-cli")
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		panic(err)
	}
}
