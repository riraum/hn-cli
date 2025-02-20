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
	// debug
	// fmt.Println(dataUnmarshalled)
	// API code below
	frontpageJSON := http.GetJSON("https://hacker-news.firebaseio.com/v0/topstories.json")
	// debug
	fmt.Println(string(frontpageJSON))

	var frontpage item.Items

	// for _, value := range frontpageJSON {
	// 	err := json.Unmarshal(value, &frontpage)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	err := json.Unmarshal(frontpageJSON, &frontpage)
	if err != nil {
		panic(err)
	}
	fmt.Println(frontpage)

	// frontpageJSONList, fpErr := item.Unmarshal(frontpageJSON)
	// if fpErr != nil {
	// 	panic(fpErr)
	// }
	// fmt.Println(frontpageJSONList)
	// debug
	// fmt.Println(string(frontpageJSONList[0]))
	// postID := frontpageJSONList
	// debug
	// fmt.Println(frontpageJSONList)
	// fmt.Println(postID)
	// postURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%v.json", postID)
	// // debug
	// fmt.Println(postURL)
	postURL := "https://hacker-news.firebaseio.com/v0/item/43094260.json"
	// debug
	// fmt.Println(postURL)
	postData := http.GetJSON(postURL)
	// debug
	// fmt.Println(string(postData))
	// var postItem item.Item
	postUnmarshalled, pErr := item.Unmarshal(postData)
	if pErr != nil {
		panic(pErr)
	}
	// debug
	fmt.Println(postUnmarshalled)
}

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
