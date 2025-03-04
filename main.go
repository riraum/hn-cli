package main

import (
	"fmt"

	"github.com/riraum/hn-cli/cmds"
	"github.com/riraum/hn-cli/format"
	"github.com/riraum/hn-cli/http"
	"github.com/riraum/hn-cli/io"
	"github.com/riraum/hn-cli/item"
	"github.com/riraum/hn-cli/ui"
)

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

	fmt.Println("postUnmarsh", postUnmarsh)

	// UI
	var input []string
	if input, err := ui.UI(); err != nil && len(input) > 1 {
		panic(err)
	}

	fmt.Println("Input slice:", input)

	var rv string

	if rv, err = cmds.Cmds(input, postUnmarsh); err != nil {
		panic(err)
	}

	fmt.Println(rv)
}
