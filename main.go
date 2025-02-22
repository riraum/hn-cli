package main

import (
	"encoding/json"
	"fmt"

	"github.com/riraum/hn-cli/http"
	"github.com/riraum/hn-cli/item"
)

func main() {
	fmt.Println("Hello hn-cli user")
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

	var timeConvert item.Item
	// set initial time as int64
	timeConvert.UnixPostTime = 1494505756
	// convert unix time stampt to time.Time
	// AbsolutePostTime := time.Unix(timeConvert.UnixPostTime, 0)
	// // debug
	// fmt.Println(AbsolutePostTime)
	// // time.time to time.Duration conversion
	// TimeSincePosting := time.Since(timeConvert.AbsolutePostTime)
	// fmt.Println(TimeSincePosting)
	// time.Duration to string conversion
	timeConvert.Time = timeConvert.RelativeTime()
	fmt.Println(timeConvert)

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

		// postUnmarshalled.AbsolutePostTime = time.Unix(postUnmarshalled.UnixPostTime, 0)
		// postUnmarshalled.TimeSincePosting = time.Since(postUnmarshalled.AbsolutePostTime)
		postUnmarshalled.Time = postUnmarshalled.RelativeTime()
		fmt.Println(postUnmarshalled)
	}
}
