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

/*
compare current time and posting time
if difference of posting time and current time is <= 1h return 1h
Figure out how to get 1hr, define it
*/
func (t Item) relativeTime() string {
	postTime := t.time
	now := time.Now()
	var returnValue string
	h, _ := time.ParseDuration("1h")
	d, _ := time.ParseDuration("1d")
	m, _ := time.ParseDuration("1m")
	y, _ := time.ParseDuration("1y")
	postAge := now.Sub(postTime)
	// if now-t.time < h {
	// 	returnValue = "<1h"
	// }
	return returnValue
}
