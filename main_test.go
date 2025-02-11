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
				title:  "Random title",
				author: "Mr Crabs",
				score:  time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix(),
				URL:    "example.com/404",
			},
			want: "1",
		},
	}

	for _, test := range tests {

	}

}
