package item

import (
	"encoding/json"
	"fmt"
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
	HoursSincePosting float64
	FormattedTime     string
	URL               string `json:"url"`
	Score             int    `json:"score"`
}

// ...
type Items []Item

var err error

func Marshall(input Item) ([]byte, error) {
	var dataMarshalled []byte

	if dataMarshalled, err = json.Marshal(input); err != nil {
		return dataMarshalled, fmt.Errorf("Failed to Marshall %w", err)
	}

	return dataMarshalled, nil
}

func Unmarshal(input []byte) (Item, error) {
	var dataUnmarshalled Item

	if err := json.Unmarshal(input, &dataUnmarshalled); err != nil {
		return dataUnmarshalled, fmt.Errorf("Failed to Unmarshall %w", err)
	}

	return dataUnmarshalled, nil
}

// time.time to time.Duration conversion
func (t Item) AddHoursSincePosting() float64 {
	return float64((time.Duration(time.Since(time.Unix(t.UnixPostTime, 0))).Hours()))
}

// ...
func (t Item) RelativeTime() string {
	elapsedHours := t.HoursSincePosting

	if elapsedHours < 1 {
		return "<1h"
	}

	if elapsedHours < hoursInADay {
		return fmt.Sprint(int64(elapsedHours), "h")
	}

	if elapsedHours < hoursInAMonth {
		return fmt.Sprint(int64(elapsedHours/hoursInADay), "d")
	}

	if elapsedHours < hoursInAYear {
		return fmt.Sprint(int64(elapsedHours/hoursInAMonth), "m")
	}

	return fmt.Sprint(int64(elapsedHours/hoursInAYear), "y")
}
