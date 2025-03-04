package cmds

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/browser"
	"github.com/riraum/hn-cli/item"
)

func OpenLink(URL string) error {
	if err := browser.OpenURL(URL); err != nil {
		return fmt.Errorf("Failed to open `%s`: %w", URL, err)
	}

	return nil
}

func Cmds(input []string, post item.Item) (string, error) {
	const hasIndex = 2

	cmd := input[0]

	var inputInt int

	var err error

	if len(input) >= hasIndex {
		if inputInt, err = strconv.Atoi(input[1]); err != nil {
			panic(err)
		}
	}

	fmt.Println("InputInt:", input)

	fmt.Println("inputInt:", inputInt)
	// To use once post print code is in function
	if cmd == "start" {
		return fmt.Sprintln("PLACEHOLDER"), err
	}
	// List commands
	if cmd == "help" {
		return fmt.Sprintln(
			"'start': Display posts\n",
			"'next': gets the next page of items\n",
			"'open X': opens the item with index/id X in the browser\n",
			"'quit': quits the program\n",
			"'refresh': reload the top items\n", "'comments': open the comments page in the browser",
		), err
	}
	// Open comments cmd
	if cmd == "comments" {
		// frontpageID := frontpageIDs[inputInt]
		// commentURL := fmt.Sprintf("https://news.ycombinator.com/item?id=%v", frontpageID)
		if err := OpenLink(post.CommentURL); err != nil {
			return post.CommentURL, fmt.Errorf("Failed to open CommentURL %w", err)
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
		if err := OpenLink(post.ArticleURL); err != nil {
			return post.ArticleURL, fmt.Errorf("Failed to open ArticleURL %w", err)
		}
	}
	// Quit command
	if cmd == "quit" {
		os.Exit(0)
	}

	return "", nil
}
