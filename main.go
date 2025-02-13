package main

import (
	"fmt"
	"math"
	"time"
)

const DAY = 24
const MONTH = 730
const YEAR = 8760

type Item struct {
	// *To use later
	// title  string
	// author string
	// score  int
	time time.Time
	URL  string
}

type Items []Item

func main() {
	fmt.Println("Hello hn-cli")
}

func (t Item) relativeTime() string {
	postTime := t.time
	postAge := time.Since(postTime)
	elapsedTime, _ := time.ParseDuration(postAge.String())
	elapsedHours := elapsedTime.Hours()
	days := elapsedHours / DAY
	months := elapsedHours / MONTH
	years := elapsedHours / YEAR

	// <  1hr
	if elapsedHours < 1 {
		return "<1h"
	}
	// number of hours, <24
	if elapsedHours < DAY {
		return fmt.Sprint(math.Round(elapsedHours), "hours")
	}
	// number of days, <30
	if elapsedHours < MONTH {
		return fmt.Sprint(math.Round(days), "d")
	}
	// number of months, <12
	if elapsedHours < YEAR {
		return fmt.Sprint(math.Round(months), "m")
	}
	// number of years
	return fmt.Sprint(math.Round(years), "y")
}
