package main

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
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
	ID     int    `json:"id"`
	// time   time.Time
	// absolutePostTime  time.Time
	timeSincePosting time.Duration
	URL              string `json:"url"`
}

type Items []Item

// type HNItem struct {
// 	id          int
// 	by          string
// 	descendants json.Number
// 	kids        []json.Number
// 	score       json.Number
// 	time        time.Time
// 	url         string
// 	title       string
// }

// type HNItems []HNItem

var HItem Item
var HItems Items

func main() {
	fmt.Println("Hello hn-cli")
}

func Unmarshal(input string) Item {
	stringBytes := []byte(input)
	var output Item
	unmarshallErr := json.Unmarshal(stringBytes, &output)
	if unmarshallErr != nil {
		panic(unmarshallErr)
	}
	// debug
	fmt.Println(output)
	return output
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
