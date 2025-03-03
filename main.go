package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"

	"github.com/pkg/browser"
	"github.com/riraum/hn-cli/http"
	"github.com/riraum/hn-cli/io"
	"github.com/riraum/hn-cli/item"
	"github.com/riraum/hn-cli/ui"
)

func openLink(URL string) error {
	return browser.OpenURL(URL)
}

func main() {
	var err error

	fmt.Println("Hello hn-cli user")
	// Get terminal size
	var tWidth int

	if tWidth, _, err = io.TermSize(); err != nil {
		panic(err)
	}

	// fmt.Println("termWidth:", tWidth)
	// API
	var frontpageJSON []byte

	if frontpageJSON, err = http.GetJSON("https://hacker-news.firebaseio.com/v0/topstories.json"); err != nil {
		panic(err)
	}

	var frontpageIDs []int

	if err = json.Unmarshal(frontpageJSON, &frontpageIDs); err != nil {
		panic(err)
	}

	for i := 0; i <= 80; i++ {
		postID := frontpageIDs[i]
		postURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", postID)

		var postData []byte

		if postData, err = http.GetJSON(postURL); err != nil {
			panic(err)
		}

		var postUnmarsh item.Item

		if postUnmarsh, err = item.Unmarshal(postData); err != nil {
			panic(err)
		}

		// Check for Ask/Show HN posts, without external URL
		if postUnmarsh.ArticleURL == "" {
			frontpageID := frontpageIDs[i]
			commentURL := fmt.Sprintf("https://news.ycombinator.com/item?id=%v", frontpageID)

			postUnmarsh.ArticleURL = commentURL
		}

		// Get CommentURL
		frontpageID := frontpageIDs[i]
		postUnmarsh.CommentURL = fmt.Sprintf("https://news.ycombinator.com/item?id=%v", frontpageID)

		// Get ArticleURL
		// postID := frontpageIDs[i]
		// postURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", postID)
		// postData := http.GetJSON(postURL)

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

	// UI
	const hasIndex = 2

	var input []string
	if input, err = ui.UI(); err != nil && len(input) > 1 {
		panic(err)
	}

	cmd := input[0]

	var inputInt int

	if len(input) >= hasIndex {
		if inputInt, err = strconv.Atoi(input[1]); err != nil {
			panic(err)
		}
	}

	// To use once post print code is in function
	if cmd == "start" {
		fmt.Println("PLACEHOLDER")
	}
	// List commands
	if cmd == "help" {
		fmt.Println(
			"'start': Display posts\n",
			"'next': gets the next page of items\n",
			"'open X': opens the item with index/id X in the browser\n",
			"'quit': quits the program\n",
			"'refresh': reload the top items\n", "'comments': open the comments page in the browser",
		)
	}
	// Open comments cmd
	if cmd == "comments" {
		frontpageID := frontpageIDs[inputInt]
		commentURL := fmt.Sprintf("https://news.ycombinator.com/item?id=%v", frontpageID)

		if err := openLink(commentURL); err != nil {
			panic(err)
		}
	}
	// Open article URL
	if cmd == "open" {
		postID := frontpageIDs[inputInt]
		postURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", postID)

		var postData []byte

		if postData, err = http.GetJSON(postURL); err != nil {
			panic(err)
		}

		postUnmarsh, err := item.Unmarshal(postData)
		if err != nil {
			panic(err)
		}

		if err := openLink(postUnmarsh.ArticleURL); err != nil {
			panic(err)
		}
	}
	// Quit command
	if cmd == "quit" {
		os.Exit(0)
	}

}
