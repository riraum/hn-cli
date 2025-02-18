package main

import (
	"bytes"
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
				timeSincePosting: 15966000000000000,
				// URL:    "example.com/404",
			},
			want: "6m",
		},
		{
			item: Item{
				timeSincePosting: 3783600000000000,
			},
			want: "1m",
		},
		{
			item: Item{
				timeSincePosting: 1105200000000000,
			},
			want: "13d",
		},
		{
			item: Item{
				timeSincePosting: 32727599999999996,
			},
			want: "1y",
		},
	}

	for _, test := range tests {
		got := test.item.relativeTime()
		if got != test.want {
			t.Errorf("Got: %v, want: %v", got, test.want)
		}
	}
}

func TestMarshall(t *testing.T) {
	tests := []struct {
		dataToMarshall Item
		want           []byte
	}{
		{
			dataToMarshall: Item{
				"Alice in Wonderland",
				"Lewis Carroll",
			},
			want: []byte(`{"title":"Alice in Wonderland","by":"Lewis Carroll"}`),
		},
	}

	for _, test := range tests {
		got, err := Marshall(test.dataToMarshall)
		if err != nil {
			panic(err)
		}

		if !bytes.Equal(got, test.want) {
			t.Errorf("Got: %v, want: %v", got, test.want)
		}
	}
}

func TestUnmarshall(t *testing.T) {
	tests := []struct {
		dataToUnmarshall []byte
		want             Item
	}{
		{
			dataToUnmarshall: []byte(`{"title":"Alice in Wonderland","by":"Lewis Carroll"}`),
			want: Item{
				"Alice in Wonderland",
				"Lewis Carroll",
			},
		},
	}

	for _, test := range tests {
		got, err := Unmarshal(test.dataToUnmarshall)
		if err != nil {
			panic(err)
		}

		if got != test.want {
			t.Errorf("Got: %v, want: %v", got, test.want)
		}
	}
}
