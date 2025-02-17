package main

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)

const (
	hoursInADay   = 24
	hoursInAMonth = 730
	hoursInAYear  = 8760
)

type Item struct {
	Title  string `json:"title"`
	Author string `json:"by"`
	Score  int    `json:"score"`
	Id     int    `json:"id"`
	// time   time.Time
	// absolutePostTime  time.Time
	timeSincePosting time.Duration
	URL              string `json:"url"`
}

type Items []Item

// type HNItem struct {
// 	id          int
// 	by          string
// 	descendants json.Number
// 	kids        []json.Number
// 	score       json.Number
// 	time        time.Time
// 	url         string
// 	title       string
// }

// type HNItems []HNItem

var HItem Item
var HItems Items

func main() {
	fmt.Println("Hello hn-cli")

	// testJSON := Item{"Alice in Wonderland", "Bob", 666, 1337, 15966000000000000, "example.com/404"}

	// testPost, testErr := json.Marshal(Item{
	// 	// "by" : "dhouston",
	// 	// "descendants" : 71,
	// 	id: 8863,
	// 	// "score" : 111,
	// 	// "time" : 1175714200,
	// 	// "title" : "My YC app: Dropbox - Throw away your USB drive",
	// 	// "type" : "story",
	// 	// "url" : "http://www.getdropbox.com/u/2/screencast.html"
	// })

	// if testErr != nil {
	// 	panic(testErr)
	// }

	// fmt.Println(string(testPost))

	// // uErr := json.Unmarshal(testPost, &HItem)
	// // if uErr != nil {
	// // 	panic(uErr)
	// // }
	// // fmt.Println(HItem)

	// testPosts, testsErr := json.Marshal(Items{Item{
	// 	// "by" : "dhouston",
	// 	// "descendants" : 71,
	// 	id: 8863,
	// 	// "score" : 111,
	// 	// "time" : 1175714200,
	// 	// "title" : "My YC app: Dropbox - Throw away your USB drive",
	// 	// "type" : "story",
	// 	// "url" : "http://www.getdropbox.com/u/2/screencast.html"
	// },
	// 	Item{
	// 		// "by" : "dhouston",
	// 		// "descendants" : 71,
	// 		id: 8865,
	// 		// "score" : 111,
	// 		// "time" : 1175714200,
	// 		// "title" : "My YC app: Dropbox - Throw away your USB drive",
	// 		// "type" : "story",
	// 		// "url" : "http://www.getdropbox.com/u/2/screencast.html"
	// 	}})
	// if testsErr != nil {
	// 	panic(testsErr)
	// }
	// fmt.Println(string(testPosts))

	testJSONPost := []byte(`{"id":8863}`)
	errPost := json.Unmarshal(testJSONPost, &HItem)
	if errPost != nil {
		panic(errPost)
	}
	fmt.Println(HItem)

	testJSONPosts := []byte(`[{
		"id": 8863},{"id": 8865}]`)
	errPosts := json.Unmarshal(testJSONPosts, &HItems)
	if errPosts != nil {
		panic(errPosts)
	}
	fmt.Println(HItems)

	// var storeData Item
	// marshallData, mErr := json.Marshal(testJSON)
	// if mErr != nil {
	// 	fmt.Println(mErr)
	// }
	// fmt.Println(string(marshallData))

	// unmarschallErr := json.Unmarshal(marshallData, &storeData)
	// if unmarschallErr != nil {
	// 	fmt.Println(unmarschallErr)
	// }
	// fmt.Println(storeData)

}

//	func (t Item) addHoursSincePosting() time.Duration {
//		t.hoursSincePosting = time.Since(t.absolutePostTime)
//		return t.hoursSincePosting
//	}
func (t Item) relativeTime() string {
	elapsedHours := t.timeSincePosting.Hours()

	if elapsedHours < 1 {
		return "<1h"
	}

	if elapsedHours < hoursInADay {
		return fmt.Sprint(math.Round(elapsedHours), "h")
	}

	if elapsedHours < hoursInAMonth {
		return fmt.Sprint(math.Round(elapsedHours/hoursInADay), "d")
	}

	if elapsedHours < hoursInAYear {
		return fmt.Sprint(math.Round(elapsedHours/hoursInAMonth), "m")
	}

	return fmt.Sprint(math.Round(elapsedHours/hoursInAYear), "y")
}
