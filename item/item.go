/*
Package items provides: stores the current posts internally

the exposed methods might be:
	`func (p Items) Get(page int) error`
	`func (p Items) Refresh() error`
	`func (p Items) Print(tWidth int) error`
	...
*/

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
	Title             string `json:"title"`
	Author            string `json:"by"`
	UnixPostTime      int64  `json:"time"`
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
		return dataMarshalled, fmt.Errorf("Failed to Marshall %w", err)
	}

	return dataMarshalled, nil
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
