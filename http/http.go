package http

import (
	"io"
	"log"
	"net/http"
)

func GetJSON(URL string) ([]byte, error) {
	var body []byte

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if body, err := io.ReadAll(resp.Body); err != nil {
		return body, err
	}

	return body, nil
}
