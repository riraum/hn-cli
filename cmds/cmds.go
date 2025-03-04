package cmds

import (
	"fmt"
	"os"

	"github.com/pkg/browser"
)

func openLink(URL string) error {
	if err := browser.OpenURL(URL); err != nil {
		return fmt.Errorf("Failed to open `%s`: %w", URL, err)
	}

	return nil
}

func Cmds(cmd string) (string, error) {
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
		// frontpageID := frontpageIDs[inputInt]
		// commentURL := fmt.Sprintf("https://news.ycombinator.com/item?id=%v", frontpageID)
		if err := openLink(postUnmarsh.CommentURL); err != nil {
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
		if err := openLink(postUnmarsh.ArticleURL); err != nil {
			panic(err)
		}
	}
	// Quit command
	if cmd == "quit" {
		os.Exit(0)
	}
}
