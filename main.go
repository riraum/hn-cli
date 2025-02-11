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

type Items []Item

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
	// now := time.Now()
	// var returnValue string
	// h, _ := time.ParseDuration("1h")
	// d, _ := time.ParseDuration("1d")
	// m, _ := time.ParseDuration("1m")
	// y, _ := time.ParseDuration("1y")
	postAge := time.Since(postTime)
	elapsedTime, _ := time.ParseDuration(postAge.String())
	elapsedHours := elapsedTime.Hours()
	days := elapsedHours / 24
	months := days * 30
	years := days * 365

	// <  1hr
	if elapsedHours < 1 {
		return "<1h"
	}
	// number of hours, <24
	if elapsedHours < 24 {
		timeMarker, h := "%f h", elapsedHours
		return fmt.Sprint(h) + timeMarker /* ... formatting with "Nh" ... */
	}
	// number of days, <30
	if elapsedHours < 720 {
		timeMarker, d := "%f d", days
		return fmt.Sprint(d) + timeMarker /* ... formatting with "Nd" ... */
	}
	// number of months, <12
	if elapsedHours < 8640 {
		timeMarker, m := "%f m", months
		return fmt.Sprint(m) + timeMarker /* ... formatting with "Nm" ... */
	}
	// number of years
	// if elapsedHours > 8640 {
	// 	return /* ... formatting with "Ny" ... */
	// }
	timeMarker, y := "%f y", years
	return fmt.Sprint(y) + timeMarker /* ... formatting with "Ny" ... */
}
