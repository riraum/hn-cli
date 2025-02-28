package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

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
	// Marshall/Unmarshall test code
	// dataToMarshall := item.Item{Title: "Alice in Wonderland", Author: "Lewis Carroll"}

	// dataMarshalled, mErr := item.Marshall(dataToMarshall)
	// if mErr != nil {
	// 	panic(mErr)
	// }
	// // debug
	// fmt.Println(string(dataMarshalled))

	// dataUnmarshalled, uErr := item.Unmarshal(dataMarshalled)
	// if uErr != nil {
	// 	panic(uErr)
	// }
	// // debug
	// fmt.Println(dataUnmarshalled)
	// Time conversion test code
	var timeConvert item.Item
	// set initial time as int64
	timeConvert.UnixPostTime = 1494505756
	timeConvert.HoursSincePosting = timeConvert.AddHoursSincePosting()
	timeConvert.FormattedTime = timeConvert.RelativeTime()
	fmt.Println(timeConvert)
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
	// debug
	// fmt.Println(string(frontpageJSON))

	var frontpageIDs []int

	err := json.Unmarshal(frontpageJSON, &frontpageIDs)
	if err != nil {
		err := fmt.Errorf("Error message during unmarshalling %g", err)
		panic(err)
	}
	// debug
	// fmt.Println(frontpageIDs)

	for i := 0; i <= 10; i++ {
		postID := frontpageIDs[i]
		postURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", postID)
		postData := http.GetJSON(postURL)

		postUnmarsh, err := item.Unmarshal(postData)
		if err != nil {
			panic(err)
		}

		postUnmarsh.Title = fmt.Sprintf("%.25s...", postUnmarsh.Title)
		postUnmarsh.HoursSincePosting = postUnmarsh.AddHoursSincePosting()
		postUnmarsh.FormattedTime = postUnmarsh.RelativeTime()

		fmt.Println(i, postUnmarsh.Score, postUnmarsh.Author, postUnmarsh.Title, postUnmarsh.FormattedTime, "ago")
	}

	// UI test code

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
