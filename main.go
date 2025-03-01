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
	fmt.Println("Hello hn-cli user")
	// Get terminal size
	var tWidth int

	tWidth, _, tErr := io.TermSize()
	if tErr != nil {
		panic(tErr)
	}

	fmt.Println("termWidth:", tWidth)
	// API
	frontpageJSON := http.GetJSON("https://hacker-news.firebaseio.com/v0/topstories.json")

	var frontpageIDs []int

	err := json.Unmarshal(frontpageJSON, &frontpageIDs)
	if err != nil {
		err := fmt.Errorf("Error message during unmarshalling %g", err)
		panic(err)
	}

	for i := 0; i <= 10; i++ {
		postID := frontpageIDs[i]
		postURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", postID)
		postData := http.GetJSON(postURL)

		postUnmarsh, err := item.Unmarshal(postData)
		if err != nil {
			panic(err)
		}

		postUnmarsh.HoursSincePosting = postUnmarsh.AddHoursSincePosting()
		postUnmarsh.FormattedTime = postUnmarsh.RelativeTime()

		// Shortening
		index := strconv.Itoa(i)
		o := fmt.Sprintln(index, postUnmarsh.Score, postUnmarsh.Author, postUnmarsh.Title, postUnmarsh.FormattedTime, "ago")

		totalLen := utf8.RuneCountInString(o)

		if totalLen > tWidth {
			const dotLen = 3

			titleLen := utf8.RuneCountInString(postUnmarsh.Title)
			// fmt.Println("titleLen:", titleLen)
			// fmt.Println("totalLen:", totalLen)
			nonReducableLen := totalLen - titleLen
			// fmt.Println("nonReducableLen:", nonReducableLen)
			reducableLen := totalLen - nonReducableLen
			// fmt.Println("reducableLen", reducableLen)

			if reducableLen > nonReducableLen {
				toReduceLen := (totalLen - tWidth)
				// fmt.Println("toReduceLen", toReduceLen)
				reducedTitleLen := (titleLen - toReduceLen - dotLen)
				// fmt.Println("reducedTitleLen", reducedTitleLen)

				postUnmarsh.Title = fmt.Sprintf("%.*s...", reducedTitleLen, postUnmarsh.Title)
			}
		}

		fo := fmt.Sprintln(index, postUnmarsh.Score, postUnmarsh.Author, postUnmarsh.Title, postUnmarsh.FormattedTime, "ago")

		fmt.Print(fo)
	}

	// UI
	const hasIndex = 2

	input, err := ui.UI()
	if len(input) > 1 && err != nil {
		panic(err)
	}

	cmd := input[0]

	var inputInt int

	if len(input) >= hasIndex {
		var err error

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
		postData := http.GetJSON(postURL)

		postUnmarsh, err := item.Unmarshal(postData)
		if err != nil {
			panic(err)
		}

		// Check for Ask/Show HN posts, without external URL
		if postUnmarsh.URL == "" {
			frontpageID := frontpageIDs[inputInt]
			commentURL := fmt.Sprintf("https://news.ycombinator.com/item?id=%v", frontpageID)

			if err := openLink(commentURL); err != nil {
				panic(err)
			}
		}

		if err := openLink(postUnmarsh.URL); err != nil {
			panic(err)
		}

		// Quit command
		if cmd == "quit" {
			os.Exit(0)
		}
	}
}
