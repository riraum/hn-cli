package main

import (
	"fmt"
	"math"
	"time"
)

const hoursInADay = 24
const hoursInAMonth = 730
const hoursInAYear = 8760

type Item struct {
	// *To use later
	// title  string
	// author string
	// score  int
	absolutePostTime  time.Time
	hoursSincePosting time.Duration
	URL               string
}

type Items []Item

func main() {
	fmt.Println("Hello hn-cli")
}

func (t Item) relativeTime() string {
	t.hoursSincePosting = time.Since(t.absolutePostTime)
	elapsedHours := t.hoursSincePosting.Hours()

	if elapsedHours < 1 {
		return "<1h"
	}

	if elapsedHours < hoursInADay {
		return fmt.Sprint(math.Round(elapsedHours), "hours")
	}

	if elapsedHours < hoursInAMonth {
		return fmt.Sprint(math.Round(elapsedHours/hoursInADay), "d")
	}

	if elapsedHours < hoursInAYear {
		return fmt.Sprint(math.Round(elapsedHours/hoursInAMonth), "m")
	}

	return fmt.Sprint(math.Round(elapsedHours/hoursInAYear), "y")
}
