package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"slices"
	"testing"

	"github.com/riraum/hn-cli/item"
)

func TestGetJSON(t *testing.T) {
	t.Run("Valid []int", func(t *testing.T) {
		t.Parallel()

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `[0,1]`)
		}))
		defer ts.Close()

		var got []int

		want := []int{0, 1}

		err := GetJSON(ts.URL, &got)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if !slices.Equal(want, got) {
			t.Fatalf("Expected %v, got %v", want, got)
		}
	},
	)
	t.Run("Valid Item", func(t *testing.T) {
		t.Parallel()

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{
  "by" : "dhouston",
  "id" : 8863,
  "score" : 111,
  "time" : 1175714200,
  "title" : "My YC app: Dropbox - Throw away your USB drive",
  "url" : "http://www.getdropbox.com/u/2/screencast.html"
}`)
		}))
		defer ts.Close()

		var got item.Item

		want := item.Item{Author: "dhouston",
			Score:        111,
			UnixPostTime: 1175714200,
			Title:        "My YC app: Dropbox - Throw away your USB drive",
			ArticleURL:   "http://www.getdropbox.com/u/2/screencast.html"}

		err := GetJSON(ts.URL, &got)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("Expected %v, got %v", want, got)
		}
	})
}
