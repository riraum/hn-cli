package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJSON(t *testing.T) {
	t.Run("ABC", func(t *testing.T) {
		t.Parallel()

		// testserver
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, []byte{0, 1})
		}))
		defer ts.Close()

		res, err := http.Get(ts.URL)
		if err != nil {
			log.Fatal(err)
		}
		intslice, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", intslice)
		// testserver end

		testString1 := struct{ ABC string }{ABC: "URL"}
		want1 := []int{0, 1}

		err = GetJSON(testString1, &want1)
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
