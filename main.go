package main

import (
	"bytes"
	"fmt"

	"github.com/riraum/hn-cli/item"

	"github.com/riraum/hn-cli/http"
)

func main() {
	fmt.Println("Hello hn-cli user")

	dataToMarshall := item.Item{Title: "Alice in Wonderland", Author: "Lewis Carroll"}

	dataMarshalled, mErr := item.Marshall(dataToMarshall)
	if mErr != nil {
		panic(mErr)
	}
	// debug
	fmt.Println(string(dataMarshalled))

	dataUnmarshalled, uErr := item.Unmarshal(dataMarshalled)
	if uErr != nil {
		panic(uErr)
	}
	// debug
	fmt.Println(dataUnmarshalled)

	frontpageJSON := http.GetJSON("https://hacker-news.firebaseio.com/v0/topstories.json")

	// debug
	fmt.Println(string(frontpageJSON))

	frontpageJSONList := bytes.Split((frontpageJSON), []byte(","))
	// debug
	fmt.Println(string(frontpageJSONList[0]))
	postID := string(frontpageJSONList[0])
	// debug
	fmt.Println(postID)

	postURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", postID)

	postData := http.GetJSON(postURL)
	fmt.Println(string(postData))
	//	for _, value := range frontpageJSON {
	//		post, uFErr := item.Unmarshal(frontpageJSON)
	//		if uFErr != nil {
	//			panic(uFErr)
	//		}
	//		frontpage = post.Title
	//	}
	//
	// fmt.Println(post)
	//
	//	for _, value := range frontpageJSON {
	//		post, uFErr := item.Unmarshal(frontpageJSON)
	//		if uFErr != nil {
	//			panic(uFErr)
	//		}
	//		fmt.Println(post)
	//		frontpage = post
	//	}
	//
	// fmt.Println(frontpage)
	// fmt.Println((unmarshallFrontpage))
}
