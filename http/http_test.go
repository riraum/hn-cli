package http

import (
	"fmt"
	"testing"

	"github.com/riraum/hn-cli/item"
)

func TestGetJSON(t *testing.T) {
	t.Run("ABC", func(t *testing.T) {
		t.Parallel()
		testString1 := fmt.Sprintf("https://news.ycombinator.com/item?id=%v", 0001)
		want1 := item.Item{}

		testString2 := "https://hacker-news.firebaseio.com/v0/topstories.json"
		want2 := []int{}
	},
	)
}
