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
				title:  "Random title",
				author: "Mr Crabs",
				score:  0,
				URL:    "example.com/404",
			},
			want: "1",
		},
	}

	for _, test := range tests {
		got := test.item.relativeTime()

		if got != test.want {
			t.Errorf("relativeTime: %v, want: %v", got, test.want)
		}
	}

}
