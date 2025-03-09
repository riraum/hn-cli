package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

func TestGetJSON(t *testing.T) {
	t.Run("ABC", func(t *testing.T) {
		t.Parallel()

		// testserver
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"":"0","":"1"}`)
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

		var want []int

		err = GetJSON(ts.URL, &want)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if !slices.Equal(want, []int{0, 1}) {
			t.Fatal("Expected want")
		}
	},
	)
}
