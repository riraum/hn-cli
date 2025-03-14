/*
Package http provides: http interactions with the HackerNews API.

It does all the GetIds, GetPosts, ...
it is a very low-level package that does only HTTP calls and maybe JSON unmarshalling.
*/

package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/riraum/hn-cli/item"
)

func GetJSON(URL string, out any) error {
	resp, err := http.Get(URL)
	if err != nil {
		return fmt.Errorf("Failed to GET `%s`: %w", URL, err)
	}

	defer resp.Body.Close()

	if out != nil {
		if err = json.NewDecoder(resp.Body).Decode(out); err != nil {
			return fmt.Errorf("Failed to decode %w", err)
		}
	}

	return nil
}

func GetPostsFromIDs(frontpageIDs []int) (item.Items, error) {
	var postUnmarshSlice item.Items

	for i := 0; i <= 10; i++ {
		var postUnmarsh item.Item

		postID := frontpageIDs[i]

		postURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", postID)

		err := GetJSON(postURL, &postUnmarsh)
		if err != nil {
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
