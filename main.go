package main

import (
	"fmt"
	"time"
)

type Item struct {
	title  string
	author string
	score  int
	time   time.Time
	URL    string
}

type Items struct {
	Item struct {
	}
}

func main() {
	fmt.Println("Hello hn-cli")
}
