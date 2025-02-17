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
	// time   time.Time
	// absolutePostTime  time.Time
	timeSincePosting time.Duration
	URL              string `json:"url"`
}

type Items []Item

type HNItem struct {
	id          json.Number
	by          string
	descendants json.Number
	kids        []json.Number
	score       json.Number
	time        time.Time
	url         string
	title       string
}

type HNItems []HNItem

var HItem Item
var HItems Items
var mf []interface{}

var HNPost HNItem
var HNFrontpage HNItems

var f interface{}
var fs []interface{}

// var dat []interface{}

// var test Item

func main() {
	fmt.Println("Hello hn-cli")

	// item1 := Items{}
	testJSON := Item{"Alice in Wonderland", "Bob", 666, 15966000000000000, "example.com/404"}

	var storeData Item
	marshallData, mErr := json.Marshal(testJSON)
	if mErr != nil {
		fmt.Println(mErr)
	}
	fmt.Println(string(marshallData))

	unmarschallErr := json.Unmarshal(marshallData, &storeData)
	if unmarschallErr != nil {
		fmt.Println(unmarschallErr)
	}
	// marshallData == []byte(`{"Title":"Alice in Wonderland","Author":"Bob","TimeSincePosting":15966000000000000}`)
	fmt.Println(storeData)

	// var unmarshallData Item
	// unmarshallErr := json.Unmarshal(marshallData, &unmarshallData)
	// if unmarshallErr != nil {
	// 	fmt.Println(unmarshallErr)
	// }
	// fmt.Println(unmarshallData)
	// debug
	// fmt.Println(test)
	// testBytes := []byte(test)
	// fmt.Println(string(testBytes))

	// merr := json.Unmarshal(testBytes, &item1)
	// if merr != nil {
	// 	fmt.Println(merr)
	// }
	// fmt.Println(item1)

	// marshallTest, mtErr := json.Marshal(Item{timeSincePosting: 15966000000000000})
	// if mtErr != nil {
	// 	panic(mtErr)
	// }
	// fmt.Println(string(marshallTest))

	errj := json.Unmarshal(body, &HItems)
	if errj != nil {
		panic(errj)
	}
	// debug
	fmt.Println(HItems)
	// fmt.Println(item1.)
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
