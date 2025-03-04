package format

import (
	"fmt"
	"log"
	"strconv"
	"unicode/utf8"

	"github.com/riraum/hn-cli/http"
	"github.com/riraum/hn-cli/item"
)

func Format(frontpageIDs []int, tWidth int) (item.Item, error) {
	var postUnmarsh item.Item

	var err error

	for i := 0; i <= 10; i++ {
		postID := frontpageIDs[i]

		postURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", postID)

		var postData []byte

		if postData, err = http.GetJSON(postURL); err != nil {
			log.Fatalln("Failed to get JSON %w", err)
		}

		if postUnmarsh, err = item.Unmarshal(postData); err != nil {
			panic(err)
		}

		// Check for Ask/Show HN posts, without external URL
		if postUnmarsh.ArticleURL == "" {
			// frontpageID := frontpageIDs[i]
			postUnmarsh.ArticleURL = fmt.Sprintf("https://news.ycombinator.com/item?id=%v", postID)
		}

		// Get CommentURL
		// frontpageID := frontpageIDs[i]
		// postUnmarsh.CommentURL = fmt.Sprintf("https://news.ycombinator.com/item?id=%v", frontpageID)
		postUnmarsh.CommentURL = fmt.Sprintf("https://news.ycombinator.com/item?id=%v", postID)

		// Get ArticleURL
		// postID := frontpageIDs[i]
		// postURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", postID)
		// postData := http.GetJSON(postURL)
		// articleURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", postID)
		// postData, err = http.GetJSON(articleURL)
		postUnmarsh.ArticleURL = postURL

		postUnmarsh.Title = fmt.Sprintf("%.25s...", postUnmarsh.Title)
		postUnmarsh.HoursSincePosting = postUnmarsh.AddHoursSincePosting()
		postUnmarsh.FormattedTime = postUnmarsh.RelativeTime()

		// Trim title
		index := strconv.Itoa(i)

		titleLen := utf8.RuneCountInString(postUnmarsh.Title)
		authorLen := utf8.RuneCountInString(postUnmarsh.Author)

		// Approximate length of the rest of the values, where a smaller variation exists, maximum observerded length taken
		const otherLen = 19
		spaceForTitle := tWidth - (otherLen + authorLen)

		if titleLen > spaceForTitle {
			postUnmarsh.Title = fmt.Sprintf("%.*s...", spaceForTitle, postUnmarsh.Title)
		}

		fmt.Println(index, postUnmarsh.Score, postUnmarsh.Author, postUnmarsh.Title, postUnmarsh.FormattedTime, "ago")
	}

	fmt.Println("postUnmarsh", postUnmarsh)

	return postUnmarsh, nil
}
