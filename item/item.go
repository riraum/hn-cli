package item

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
	// Score  int    `json:"score"`
	// ID                int `json:"id"`
	absolutePostTime  time.Time
	timeSincePosting  time.Duration
	hoursSincePosting time.Duration
	// URL string `json:"url"`
}

// ...
type Items []Item

func Marshall(input Item) ([]byte, error) {
	dataMarshalled, err := json.Marshal(input)
	if err != nil {
		return dataMarshalled, err
	}

	return dataMarshalled, nil
}

func Unmarshal(input []byte) (Item, error) {
	var dataUnmarshalled Item

	err := json.Unmarshal(input, &dataUnmarshalled)
	if err != nil {
		return dataUnmarshalled, err
	}

	return dataUnmarshalled, nil
}

func (t Item) AddHoursSincePosting() time.Duration {
	t.hoursSincePosting = time.Since(t.absolutePostTime)
	return t.hoursSincePosting
}

// ...
func (t Item) RelativeTime() string {
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
