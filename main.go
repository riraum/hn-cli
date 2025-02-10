package main

import "fmt"

type Item struct {
	title  string
	author string
	score  int
	time   string
	URL    string
}

type Items struct {
	Item struct {
	}
}

func main() {
	fmt.Println("Hello hn-cli")
}
