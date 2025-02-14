package main

import (
	"testing"
	"time"
)

func TestRelativeTime(t *testing.T) {
	tests := []struct {
		item Item
		want time.Duration
	}{
		{
			item: Item{
				// title:  "Random title",
				// author: "Mr Crabs",
				// score:  0,
				absolutePostTime: time.Date(2024, 8, 13, 14, 30, 45, 100, time.UTC),
				// URL:    "example.com/404",
			},
			// want: "6m",
		},
		{
			item: Item{
				absolutePostTime: time.Date(2025, 1, 01, 14, 30, 45, 100, time.UTC),
			},
			// want: "1m",
		},
		{
			item: Item{
				absolutePostTime: time.Date(2025, 2, 01, 14, 30, 45, 100, time.UTC),
			},
			// want: "12d",
		},
		{
			item: Item{
				absolutePostTime: time.Date(2024, 2, 01, 14, 30, 45, 100, time.UTC),
			},
			// want: "1y",
		},
	}

	for _, test := range tests {
		now := time.Now()
		diff := now.Sub(test.item.absolutePostTime)
		test.want = diff

		test.item.addHoursSincePosting()
		// absoluteTime := Item{absolutePostTime: test.item.absolutePostTime}
		got := test.item.hoursSincePosting

		// absoluteTime := Item{absolutePostTime: test.item.absolutePostTime}
		// got := absolutePostTime.relativeTime()

		if got != test.want {
			t.Errorf("Got: %v, want: %v", got, test.want)
		}
	}
}
