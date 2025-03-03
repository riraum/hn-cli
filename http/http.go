package http

import (
	"fmt"
	"io"
	"net/http"
)

func GetJSON(URL string) ([]byte, error) {
	var body []byte

	resp, err := http.Get(URL)
	if err != nil {
		return nil, fmt.Errorf("Failed to GET `%s`: %w", URL, err)
	}

	defer resp.Body.Close()

	if body, err := io.ReadAll(resp.Body); err != nil {
		return body, fmt.Errorf("Failed to read response body: %w", err)
	}

	return body, nil
}
