package http

import (
	"fmt"
	"io"
	"net/http"
)

func GetJSON(URL string) ([]byte, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, fmt.Errorf("Failed to GET `%s`: %w", URL, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return body, fmt.Errorf("Failed to read response body: %w", err)
	}

	return body, nil
}
