package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/browser"
	"github.com/riraum/hn-cli/http"
	"github.com/riraum/hn-cli/io"
	"github.com/riraum/hn-cli/item"
	"github.com/riraum/hn-cli/ui"
)

func main() {
	fmt.Println("Hello hn-cli user")
	// Marshall/Unmarshall test code
	// dataToMarshall := item.Item{Title: "Alice in Wonderland", Author: "Lewis Carroll"}

	// dataMarshalled, mErr := item.Marshall(dataToMarshall)
	// if mErr != nil {
	// 	panic(mErr)
	// }
	// // debug
	// fmt.Println(string(dataMarshalled))

	// dataUnmarshalled, uErr := item.Unmarshal(dataMarshalled)
	// if uErr != nil {
	// 	panic(uErr)
	// }
	// // debug
	// fmt.Println(dataUnmarshalled)
	// Time conversion test code
	var timeConvert item.Item
	// set initial time as int64
	timeConvert.UnixPostTime = 1494505756
	timeConvert.FormattedTime = timeConvert.RelativeTime()
	fmt.Println(timeConvert)
	// Get terminal size test code
	var tWidth int

	var tHeight int

	tWidth, tHeight, tErr := io.TermSize()
	if tErr != nil {
		panic(tErr)
	}

	fmt.Println("Size:", tWidth, tHeight)

	// UI test code
	var input string

	input, uErr := ui.UI()
	if uErr != nil {
		panic(uErr)
	}

	// Open article URL
	if input == "open" {
		frontpage := item.Items{item.Item{
			URL: "https://github.com",
		}}
		// inputIndex
		openURL := frontpage[0].URL

		err := browser.OpenURL(openURL)
		if err != nil {
			panic(err)
		}
	}
	// Quit command
	if input == "quit" {
		os.Exit(0)
	}

	fmt.Println(input)
	// API code below
	frontpageJSON := http.GetJSON("https://hacker-news.firebaseio.com/v0/topstories.json")
	// debug
	// fmt.Println(string(frontpageJSON))

	var frontpageIDs []int

	err := json.Unmarshal(frontpageJSON, &frontpageIDs)
	if err != nil {
		panic(err)
	}
	// debug
	// fmt.Println(frontpageIDs)
	for _, postID := range frontpageIDs {
		postURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", postID)
		postData := http.GetJSON(postURL)

		postUnmarshalled, pErr := item.Unmarshal(postData)
		if pErr != nil {
			panic(pErr)
		}

		postUnmarshalled.FormattedTime = postUnmarshalled.RelativeTime()
		fmt.Println(postUnmarshalled)
	}
}
