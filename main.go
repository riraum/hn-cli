package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/riraum/hn-cli/assemble"
	"github.com/riraum/hn-cli/cmds"
	"github.com/riraum/hn-cli/io"
	"github.com/riraum/hn-cli/item"
	"github.com/riraum/hn-cli/ui"
)

func main() {
	fmt.Println("Hello hn-cli user")

	const errTxt = "Goodbye"

	// Get terminal size
	tWidth, err := io.TermSize()
	if err != nil {
		fmt.Println(errTxt, err)
		os.Exit(1)
	}

	// API
	var posts item.Items

	posts, err = assemble.GetAndFormatPosts(tWidth)
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

	_, err = cmds.Run(input[0], posts[inputInt], tWidth)
	if err != nil {
		fmt.Println(errTxt, err)
		os.Exit(1)
	}
}
