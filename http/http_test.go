package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

func TestGetJSON(t *testing.T) {
	t.Run("Valid []int", func(t *testing.T) {
		t.Parallel()

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `[0,1]`)
		}))
		defer ts.Close()

		// res, err := http.Get(ts.URL)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// byteOut, err := io.ReadAll(res.Body)
		// res.Body.Close()

		// if err != nil {
		// 	log.Fatal(err)
		// }

		// fmt.Printf("%s", byteOut)

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
}
