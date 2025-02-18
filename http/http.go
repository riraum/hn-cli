package http

import (
	"io"
	"net/http"
)

func GetJSON(URL string) []byte {
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return body
}
