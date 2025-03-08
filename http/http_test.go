package http

import (
	"testing"
)

func TestGetJSON(t *testing.T) {
	t.Run("ABC", func(t *testing.T) {
		t.Parallel()
		testString1 := struct{ ABC string }{ABC: "URL"}
		want1 := []int{0, 1}

		err := GetJSON(testString1, &want1)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if !testString1.want1 {
			t.Fatal("Expected want")
		}

		testString2 := "https://hacker-news.firebaseio.com/v0/topstories.json"
		want2 := []int{}
	},
	)
}
