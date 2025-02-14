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
				hoursSincePosting: 4433,
				// URL:    "example.com/404",
			},
			want: "6m",
		},
		{
			item: Item{
				// absolutePostTime: time.Date(2025, 1, 01, 14, 30, 45, 100, time.UTC),
				hoursSincePosting: 1049,
			},
			want: "1m",
		},
		{
			item: Item{
				// absolutePostTime: time.Date(2025, 2, 01, 14, 30, 45, 100, time.UTC),
				hoursSincePosting: 305,
			},
			want: "12d",
		},
		{
			item: Item{
				// absolutePostTime: time.Date(2024, 2, 01, 14, 30, 45, 100, time.UTC),
				hoursSincePosting: 9089,
			},
			want: "1y",
		},
	}

	for _, test := range tests {
		// now := time.Now()
		// diff := now.Sub(test.item.absolutePostTime)
		// test.want = diff

		// test.item.addHoursSincePosting()
		// absoluteTime := Item{absolutePostTime: test.item.absolutePostTime}
		// got := test.item.hoursSincePosting
		// x := test.item
		got := test.item.relativeTime()
		// absoluteTime := Item{absolutePostTime: test.item.absolutePostTime}
		// got := absolutePostTime.relativeTime()

		if got != test.want {
			t.Errorf("Got: %v, want: %v", got, test.want)
		}
	}
}
