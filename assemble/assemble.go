package assemble

import (
	"fmt"

	"github.com/riraum/hn-cli/format"
	"github.com/riraum/hn-cli/http"
	"github.com/riraum/hn-cli/item"
)

func Show(tWidth int) (item.Items, error) {
	fmt.Println("Talking to the API...")

	frontpageJSON, err := http.GetJSON("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		return nil, fmt.Errorf("Failed to GetJSON %w", err)
	}

	frontpageIDs, err := item.UnmarshallSlice(frontpageJSON)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshall slice %w", err)
	}

	posts, err := http.GetPostsFromIDs(frontpageIDs)
	if err != nil {
		return nil, fmt.Errorf("Failed to GetPostsFromIDs %w", err)
	}

	err = format.Format(posts, tWidth)
	if err != nil {
		return nil, fmt.Errorf("Failed to Format %w", err)
	}

	return posts, nil
}
