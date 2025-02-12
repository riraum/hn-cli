package main

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

type Item struct {
	title  string
	author string
	score  int
	time   time.Time
	URL    string
}

type Items []Item

func main() {
	fmt.Println("Hello hn-cli")
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

/*
compare current time and posting time
if difference of posting time and current time is <= 1h return 1h
Figure out how to get 1hr, define it
*/
func (t Item) relativeTime() string {
	postTime := t.time
	// now := time.Now()
	postAge := time.Since(postTime)
	elapsedTime, _ := time.ParseDuration(postAge.String())
	elapsedHours := elapsedTime.Hours()
	days := elapsedHours / 24
	months := elapsedHours / 730
	years := elapsedHours / 8760

	// <  1hr
	if elapsedHours < 1 {
		return "<1h"
	}
	// number of hours, <24
	if elapsedHours < 24 {
		return fmt.Sprint(math.Round(elapsedHours), "hours") /* ... formatting with "Nh" ... */
	}
	// number of days, <30
	if elapsedHours < 720 {
		return fmt.Sprint(math.Round(days), "d") /* ... formatting with "Nd" ... */
	}
	// number of months, <12
	if elapsedHours < 8640 {
		return fmt.Sprint(math.Round(months), "m") /* ... formatting with "Nm" ... */
	}
	// number of years
	return fmt.Sprint(math.Round(years), "y") /* ... formatting with "Ny" ... */
}
