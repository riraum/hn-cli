package main

import (
	"fmt"
	"math"
	"time"
)

const HOURS_IN_A_DAY = 24
const HOURS_IN_A_MONTH = 730
const HOURS_IN_A_YEAR = 8760

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

	if elapsedHours < 1 {
		return "<1h"
	}

	if elapsedHours < HOURS_IN_A_DAY {
		return fmt.Sprint(math.Round(elapsedHours), "hours")
	}

	if elapsedHours < HOURS_IN_A_MONTH {
		return fmt.Sprint(math.Round(elapsedHours/HOURS_IN_A_DAY), "d")
	}

	if elapsedHours < HOURS_IN_A_YEAR {
		return fmt.Sprint(math.Round(elapsedHours/HOURS_IN_A_MONTH), "m")
	}

	return fmt.Sprint(math.Round(elapsedHours/HOURS_IN_A_YEAR), "y")
}
