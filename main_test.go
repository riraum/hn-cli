package main

import (
	"fmt"
	"testing"
)

func TestRelativeTime(t *testing.T) {
	tests := []struct {
		item Item
		want string
	}{
		{
			item: Item{
				// title:  "Random title",
				// author: "Mr Crabs",
				// score:  0,
				// absolutePostTime: time.Date(2024, 8, 13, 14, 30, 45, 100, time.UTC),
				// hoursSincePosting: "4435h35m7.2785059s",
				// hoursSincePosting: time.ParseDuration("4435h"),
				hoursSincePosting: 15966000000000000,

				// URL:    "example.com/404",
			},
			want: "6m",
		},
		{
			item: Item{
				// absolutePostTime: time.Date(2025, 1, 01, 14, 30, 45, 100, time.UTC),
				// hoursSincePosting: 1051h35m7.2785799s,
				hoursSincePosting: 3783600000000000,
			},
			want: "1m",
		},
		{
			item: Item{
				// absolutePostTime: time.Date(2025, 2, 01, 14, 30, 45, 100, time.UTC),
				// hoursSincePosting: 307h35m7.2785829s,
				hoursSincePosting: 1105200000000000,
			},
			want: "12d",
		},
		{
			item: Item{
				// absolutePostTime: time.Date(2024, 2, 01, 14, 30, 45, 100, time.UTC),
				// hoursSincePosting: 9091h35m7.2786449s,
				hoursSincePosting: 32727599999999996,
			},
			want: "1y",
		},
	}

	for _, test := range tests {
		// duration.Round(time.Hour).String()
		// now := time.Now()
		// diff := now.Sub(test.item.absolutePostTime)
		// test.item.hoursSincePosting = diff

		// test.item.addHoursSincePosting()
		// absoluteTime := Item{absolutePostTime: test.item.absolutePostTime}
		// got := test.item.hoursSincePosting
		// x := test.item
		got := test.item.relativeTime()
		// debug
		fmt.Println(test.item.hoursSincePosting)
		// absoluteTime := Item{absolutePostTime: test.item.absolutePostTime}
		// got := absolutePostTime.relativeTime()

		if got != test.want {
			t.Errorf("Got: %v, want: %v", got, test.want)
		}
	}
}
