package main

import (
	"testing"
)

func TestRelativeTime(t *testing.T) {
	tests := []struct {
		item Item
		want string
	}{
		item: Item{
			title: "Random title",
			author: "Mr Crabs",
			score: "1337",
			time: "",
			URL: "example.com/404",
		},
		want: "1",
	},

	
	}
