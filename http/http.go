package http

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/riraum/hn-cli/item"
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

func GetPostsFromIDs(frontpageIDs []int) (item.Items, error) {
	var postUnmarshSlice item.Items

	var err error

	for i := 0; i <= 10; i++ {
		var postUnmarsh item.Item

		postID := frontpageIDs[i]

		postURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", postID)

		var postData []byte

		if postData, err = GetJSON(postURL); err != nil {
			log.Fatalln("Failed to get JSON %w", err)
		}

		if postUnmarsh, err = item.Unmarshal(postData); err != nil {
			log.Fatalln("Failed to Unmarshall %w", err)
		}

		// Check for Ask/Show HN posts, without external URL
		if postUnmarsh.ArticleURL == "" {
			postUnmarsh.ArticleURL = fmt.Sprintf("https://news.ycombinator.com/item?id=%v", postID)
		}

		// Get CommentURL
		postUnmarsh.CommentURL = fmt.Sprintf("https://news.ycombinator.com/item?id=%v", postID)

		postUnmarshSlice = append(postUnmarshSlice, postUnmarsh)
	}

	return postUnmarshSlice, nil
}
