package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"
)

const (
	hoursInADay   = 24
	hoursInAMonth = 730
	hoursInAYear  = 8760
)

type Item struct {
	// title  string
	// author string
	// score  int
	// absolutePostTime  time.Time
	timeSincePosting time.Duration
	URL              string
}

type Items []Item

func main() {
	fmt.Println("Hello hn-cli")
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	// if err,
	json.Unmarshal(body)
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
