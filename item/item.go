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
	UnixPostTime      int64 `json:"time"`
	hoursSincePosting time.Duration
	FormattedTime     string
	URL               string `json:"url"`
	Score             int    `json:"score"`
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

	dataUnmarshalled.hoursSincePosting = time.Duration(time.Since(time.Unix(dataUnmarshalled.UnixPostTime, 0)).Hours())

	return dataUnmarshalled, nil
}

// ...
func (t Item) RelativeTime() string {
	elapsedHours := t.hoursSincePosting.Hours()

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
