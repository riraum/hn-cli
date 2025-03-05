package main

import (
	"fmt"
	"strconv"

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

	postUnmarshSlice, err := format.Format(frontpageIDs, tWidth)
	if err != nil {
		panic(err)
	}

	fmt.Println("postUnmarshSlice", postUnmarshSlice)

	// UI
	input, err := ui.UI()
	if err != nil && len(input) > 1 {
		panic(err)
	}

	const hasIndex = 2

	var inputInt int

	if len(input) >= hasIndex {
		if inputInt, err = strconv.Atoi(input[1]); err != nil {
			panic(err)
		}
	}

	fmt.Println("Input slice:", input)

	err = cmds.Run(input, postUnmarshSlice[inputInt])
	if err != nil {
		panic(err)
	}
}
