package format

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/riraum/hn-cli/item"
)

func Format(posts item.Items, tWidth int) error {
	for i := 0; i <= 10; i++ {
		postUnmarsh := posts[i]

		postUnmarsh.HoursSincePosting = postUnmarsh.AddHoursSincePosting()
		postUnmarsh.FormattedTime = postUnmarsh.RelativeTime()

		// Trim title
		index := strconv.Itoa(i)

		titleLen := utf8.RuneCountInString(postUnmarsh.Title)
		authorLen := utf8.RuneCountInString(postUnmarsh.Author)

		// Approximate length of the rest of the values, where a smaller variation exists, maximum observerded length taken
		const otherLen = 19
		spaceForTitle := tWidth - (otherLen + authorLen)

		if titleLen > spaceForTitle {
			postUnmarsh.Title = fmt.Sprintf("%.*s...", spaceForTitle, postUnmarsh.Title)
		}

		fmt.Println(index, postUnmarsh.Score, postUnmarsh.Author, postUnmarsh.Title, postUnmarsh.FormattedTime, "ago")
	}

	return nil
}
