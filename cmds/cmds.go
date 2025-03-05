package cmds

import (
	"fmt"
	"os"

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

func Run(input string, post item.Item) error {
	// fmt.Println("input:", input)
	// To use once post print code is in function
	if input == "" {
		fmt.Sprintln("You didn't enter anything. Will print help", help())
	}

	switch input {
	// case "":
	// fmt.Sprintln("You didn't enter anything. Will print help", help())
	case "start":
		fmt.Sprintln("PLACEHOLDER")
	// List commands
	case "help":
		fmt.Print(help())
	// Open comments cmd
	case "comments":
		if err := openCommentURL(post.CommentURL); err != nil {
			return fmt.Errorf("Failed to open URL %w", err)
		}
	// Open article URL
	case "open":
		if err := openArticleURL(post.ArticleURL); err != nil {
			return fmt.Errorf("Failed to open URL %w", err)
		}
	// Quit command
	case "quit":
		quit()
	}

	return nil
}
