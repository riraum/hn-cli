package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/riraum/hn-cli/cmds"
	"github.com/riraum/hn-cli/format"
	"github.com/riraum/hn-cli/http"
	"github.com/riraum/hn-cli/io"
	"github.com/riraum/hn-cli/item"
	"github.com/riraum/hn-cli/ui"
)

// func openLink(URL string) error {
// 	if err := browser.OpenURL(URL); err != nil {
// 		return fmt.Errorf("Failed to open `%s`: %w", URL, err)
// 	}

// 	return nil
// }

func main() {
	fmt.Println("Hello hn-cli user")

	// Get terminal size
	tWidth, err := io.TermSize()
	if err != nil {
		panic(err)
	}

	// API
	frontpageJSON, err := http.GetJSON("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		panic(err)
	}
	// fmt.Println("frontpageJSON", frontpageJSON)

	frontpageIDs, err := item.UnmarshallSlice(frontpageJSON)
	if err != nil {
		panic(err)
	}

	// fmt.Println("frontpageIDs", frontpageIDs)

	var postUnmarsh item.Item

	postUnmarsh, err = format.Format(frontpageIDs, tWidth)
	if err != nil {
		panic(err)
	}

	// UI
	const hasIndex = 2

	var input []string
	if input, err = ui.UI(); err != nil && len(input) > 1 {
		panic(err)
	}

	fmt.Println("Input slice:", input)

	cmd := input[0]

	var inputInt int

	if len(input) >= hasIndex {
		if inputInt, err = strconv.Atoi(input[1]); err != nil {
			panic(err)
		}
	}

	fmt.Println("InputInt:", input)

	fmt.Println("inputInt:", inputInt)

	// cmd.Cmds()

	// // To use once post print code is in function

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
		// frontpageID := frontpageIDs[inputInt]
		// commentURL := fmt.Sprintf("https://news.ycombinator.com/item?id=%v", frontpageID)
		if err := cmds.OpenLink(postUnmarsh.CommentURL); err != nil {
			panic(err)
		}
	}

	// Open article URL

	if cmd == "open" {
		// postID := frontpageIDs[i]
		// postURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", postID)
		// var postData []byte
		// if postData, err = http.GetJSON(postURL); err != nil {
		// 	panic(err)
		// }
		// postUnmarsh, err := item.Unmarshal(postData)
		// if err != nil {
		// 	panic(err)
		// }
		if err := cmds.OpenLink(postUnmarsh.ArticleURL); err != nil {
			panic(err)
		}
	}

	// Quit command

	if cmd == "quit" {
		os.Exit(0)
	}
}
