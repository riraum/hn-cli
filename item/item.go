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
	ArticleURL        string `json:"url"`
	CommentURL        string
	Score             int `json:"score"`
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
