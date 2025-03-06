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

func main() {
	fmt.Println("Hello hn-cli user\nTalking to the API...")

	const errTxt = "Goodbye"

	// Get terminal size
	tWidth, err := io.TermSize()
	if err != nil {
		fmt.Println("errTxt", err)
		os.Exit(1)
	}

	// API
	frontpageJSON, err := http.GetJSON("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		fmt.Println(errTxt, err)
		os.Exit(1)
	}

	frontpageIDs, err := item.UnmarshallSlice(frontpageJSON)
	if err != nil {
		fmt.Println(errTxt, err)
		os.Exit(1)
	}

	posts, err := http.GetPostsFromIDs(frontpageIDs)
	if err != nil {
		fmt.Println(errTxt, err)
		os.Exit(1)
	}

	err = format.Format(posts, tWidth)
	if err != nil {
		fmt.Println(errTxt, err)
		os.Exit(1)
	}

	// UI
	input, err := ui.UI()
	if err != nil && len(input) > 1 {
		fmt.Println(errTxt, err)
		os.Exit(1)
	}

	const hasIndex = 2

	var inputInt int

	if len(input) >= hasIndex {
		if inputInt, err = strconv.Atoi(input[1]); err != nil {
			fmt.Println(errTxt, err)
			os.Exit(1)
		}
	}

	err = cmds.Run(input[0], posts[inputInt])
	if err != nil {
		fmt.Println(errTxt, err)
		os.Exit(1)
	}
}
