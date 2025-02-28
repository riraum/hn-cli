package main

import (
	"encoding/json"
	"fmt"
	"os"

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
	fmt.Println("Hello hn-cli user")
	// Get terminal size test code
	var tWidth int

	var tHeight int

	tWidth, tHeight, tErr := io.TermSize()
	if tErr != nil {
		panic(tErr)
	}

	fmt.Println("Size:", tWidth, tHeight)
	// API code below
	frontpageJSON := http.GetJSON("https://hacker-news.firebaseio.com/v0/topstories.json")

	var frontpageIDs []int

	err := json.Unmarshal(frontpageJSON, &frontpageIDs)
	if err != nil {
		panic(err)
	}

	const accountForRestStr = 30

	for i := 0; i <= 10; i++ {
		postID := frontpageIDs[i]
		postURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", postID)
		postData := http.GetJSON(postURL)

		postUnmarsh, pErr := item.Unmarshal(postData)
		if pErr != nil {
			panic(pErr)
		}

		postUnmarsh.Title = fmt.Sprintf("%.*s...", tWidth-accountForRestStr, postUnmarsh.Title)
		postUnmarsh.HoursSincePosting = postUnmarsh.AddHoursSincePosting()
		postUnmarsh.FormattedTime = postUnmarsh.RelativeTime()

		fmt.Println(i, postUnmarsh.Score, postUnmarsh.Author, postUnmarsh.Title, postUnmarsh.FormattedTime, "ago")
	}

	// UI
	var input string

	input, uErr := ui.UI()
	if uErr != nil {
		panic(uErr)
	}
	// To use once post print code is in function
	if input == "start" {
		fmt.Println("PLACEHOLDER")
	}
	// List commands
	if input == "help" {
		fmt.Println(
			"'start': Display posts\n",
			"'next': gets the next page of items\n",
			"'open X': opens the item with index/id X in the browser\n",
			"'quit': quits the program\n",
			"'refresh': reload the top items\n", "'comments': open the comments page in the browser",
		)
	}
	// Open comments cmd
	if input == "comments" {
		frontpageID := 8863
		commentURL := fmt.Sprintf("https://news.ycombinator.com/item?id=%v", frontpageID)

		if err := openLink(commentURL); err != nil {
			panic(err)
		}
	}
	// Open article URL
	if input == "open" {
		frontpage := item.Items{item.Item{
			URL: "https://github.com",
		}}
		// inputIndex
		openURL := frontpage[0].URL

		if err := openLink(openURL); err != nil {
			panic(err)
		}
	}
	// Quit command
	if input == "quit" {
		os.Exit(0)
	}
}
