package main

import (
	"testing"
	"time"
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
				absolutePostTime: time.Date(2024, 8, 13, 14, 30, 45, 100, time.UTC),
				// URL:    "example.com/404",
			},
			want: "6m",
		},
		{
			item: Item{
				absolutePostTime: time.Date(2025, 1, 01, 14, 30, 45, 100, time.UTC),
			},
			want: "1m",
		},
		{
			item: Item{
				absolutePostTime: time.Date(2025, 2, 01, 14, 30, 45, 100, time.UTC),
			},
			want: "12d",
		},
		{
			item: Item{
				absolutePostTime: time.Date(2024, 2, 01, 14, 30, 45, 100, time.UTC),
			},
			want: "1y",
		},
	}

	for _, test := range tests {
		absoluteTime := Item{absolutePostTime: test.item.absolutePostTime}
		got := absoluteTime.relativeTime()

		if got != test.want {
			t.Errorf("relativeTime: %v, want: %v", got, test.want)
		}
	}
}
