package main

import (
	"fmt"
	"math"
	"time"

	"github.com/riraum/hn-cli/http"
)

const (
	hoursInADay   = 24
	hoursInAMonth = 730
	hoursInAYear  = 8760
)

type Item struct {
	Title  string `json:"title"`
	Author string `json:"by"`
	Score  int    `json:"score"`
	// time   time.Time
	// absolutePostTime  time.Time
	timeSincePosting time.Duration
	URL              string `json:"url"`
}

type Items []Item

func main() {
	fmt.Println("Hello hn-cli")

	frontpageJSON := http.GetJSON("https://hacker-news.firebaseio.com/v0/topstories.json")

	// debug
	fmt.Println(string(frontpageJSON))
}

//	func (t Item) addHoursSincePosting() time.Duration {
//		t.hoursSincePosting = time.Since(t.absolutePostTime)
//		return t.hoursSincePosting
//	}
func (t Item) relativeTime() string {
	elapsedHours := t.timeSincePosting.Hours()

	if elapsedHours < 1 {
		return "<1h"
	}

	if elapsedHours < hoursInADay {
		return fmt.Sprint(math.Round(elapsedHours), "h")
	}

	if elapsedHours < hoursInAMonth {
		return fmt.Sprint(math.Round(elapsedHours/hoursInADay), "d")
	}

	if elapsedHours < hoursInAYear {
		return fmt.Sprint(math.Round(elapsedHours/hoursInAMonth), "m")
	}

	return fmt.Sprint(math.Round(elapsedHours/hoursInAYear), "y")
}
