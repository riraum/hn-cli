package main

import (
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
				timeSincePosting: 15966000000000000,

				// URL:    "example.com/404",
			},
			want: "6m",
		},
		{
			item: Item{
				// absolutePostTime: time.Date(2025, 1, 01, 14, 30, 45, 100, time.UTC),
				// hoursSincePosting: 1051h35m7.2785799s,
				timeSincePosting: 3783600000000000,
			},
			want: "1m",
		},
		{
			item: Item{
				// absolutePostTime: time.Date(2025, 2, 01, 14, 30, 45, 100, time.UTC),
				// hoursSincePosting: 307h35m7.2785829s,
				timeSincePosting: 1105200000000000,
			},
			want: "13d",
		},
		{
			item: Item{
				// absolutePostTime: time.Date(2024, 2, 01, 14, 30, 45, 100, time.UTC),
				// hoursSincePosting: 9091h35m7.2786449s,
				timeSincePosting: 32727599999999996,
			},
			want: "1y",
		},
	}

	for _, test := range tests {
		got := test.item.relativeTime()
		// debug
		// fmt.Println(test.item.timeSincePosting)
		if got != test.want {
			t.Errorf("Got: %v, want: %v", got, test.want)
		}
	}
}
