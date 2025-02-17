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
	Title  string
	Author string
	// score  int
	// absolutePostTime  time.Time
	timeSincePosting time.Duration
	// URL              string
}

type Items []Item

var m Item
var ms Items
var mf map[string]interface{}

// var test Item

func main() {
	fmt.Println("Hello hn-cli")

	// item1 := Items{}
	testJSON := Item{"Alice in Wonderland", "Bob", 15966000000000000}

	var storeData Item
	marshallData, mErr := json.Marshal(testJSON)
	if mErr != nil {
		fmt.Println(mErr)
	}
	fmt.Println(marshallData)

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

	// resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// body, _ := io.ReadAll(resp.Body)
	// // debug
	// fmt.Println(body)

	// errj := json.Unmarshal(body, &item1)
	// if errj != nil {
	// 	panic(errj)
	// }
	// debug
	// fmt.Println(mf)
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
