package main

import (
	"testing"
)

// func TestRelativeTime(t *testing.T) {
// 	tests := []struct {
// 		item Item
// 		want string
// 	}{
// 		{
// 			item: Item{
// 				// title:  "Random title",
// 				// author: "Mr Crabs",
// 				// score:  0,
// 				timeSincePosting: 15966000000000000,
// 				// URL:    "example.com/404",
// 			},
// 			want: "6m",
// 		},
// 		{
// 			item: Item{
// 				timeSincePosting: 3783600000000000,
// 			},
// 			want: "1m",
// 		},
// 		{
// 			item: Item{
// 				timeSincePosting: 1105200000000000,
// 			},
// 			want: "13d",
// 		},
// 		{
// 			item: Item{
// 				timeSincePosting: 32727599999999996,
// 			},
// 			want: "1y",
// 		},
// 	}

// 	for _, test := range tests {
// 		got := test.item.relativeTime()
// 		if got != test.want {
// 			t.Errorf("Got: %v, want: %v", got, test.want)
// 		}
// 	}
// }

func TestMarshall(t *testing.T) {
	tests := []struct {
		dataToMarshall Item
		want           string
	}{
		{
			dataToMarshall: Item{
				"Alice in Wonderland",
				"Lewis Carroll",
			},
			want: {"title": "Alice in Wonderland", "by": "Lewis Carroll"},
		},
	}

	for _, test := range tests {
		got := Marshall(test.dataToMarshall)
		if got != test.want {
			t.Errorf("Got: %v, want: %v", got, test.want)
		}
	}
}

func TestUnmarshall(t *testing.T) {
	tests := []struct {
		item string
		want Item
	}{
		{
			item: "Random title",
			want: Item{
				Title: "Random title"},
		},
	}

	for _, test := range tests {
		got := Unmarshal(test.item)
		if got != test.want {
			t.Errorf("Got: %v, want: %v", got, test.want)
		}
	}
}
