package cmds

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/browser"
	"github.com/riraum/hn-cli/item"
)

func openLink(URL string) error {
	if err := browser.OpenURL(URL); err != nil {
		return fmt.Errorf("Failed to open `%s`: %w", URL, err)
	}

	return nil
}

func openCommentURL(URL string) error {
	err := openLink(URL)
	if err != nil {
		return fmt.Errorf("Failed to open CommentURL %w", err)
	}

	return nil
}

func openArticleURL(URL string) error {
	err := openLink(URL)
	if err != nil {
		return fmt.Errorf("Failed to open ArticleURL %w", err)
	}

	return nil
}

func help() string {
	return fmt.Sprintln(
		"'start': Display posts\n",
		"'next': gets the next page of items\n",
		"'open X': opens the item with index/id X in the browser\n",
		"'quit': quits the program\n",
		"'refresh': reload the top items\n", "'comments': open the comments page in the browser",
	)
}

func quit() {
	os.Exit(0)
}

func Run(input []string, post item.Item) error {
	const hasIndex = 2

	cmd := input[0]

	var inputInt int

	var err error

	if len(input) >= hasIndex {
		if inputInt, err = strconv.Atoi(input[1]); err != nil {
			return fmt.Errorf("Failed Atoi conversion %w", err)
		}
	}

	fmt.Println("input:", input)

	fmt.Println("inputInt:", inputInt)
	// To use once post print code is in function
	switch cmd {
	case "start":
		fmt.Sprintln("PLACEHOLDER")
	// List commands
	case "help":
		fmt.Print(help())
	// Open comments cmd
	case "comments":
		// frontpageID := frontpageIDs[inputInt]
		// commentURL := fmt.Sprintf("https://news.ycombinator.com/item?id=%v", frontpageID)
		if err := openCommentURL(post.CommentURL); err != nil {
			return fmt.Errorf("Failed to open URL %w", err)
		}
	// Open article URL
	case "open":
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
		if err := openArticleURL(post.ArticleURL); err != nil {
			return fmt.Errorf("Failed to open URL %w", err)
		}
	// Quit command
	case "quit":
		quit()
	}

	return nil
}
