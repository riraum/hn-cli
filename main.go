package main

import (
	"encoding/json"
	"fmt"
)

// const (
// 	hoursInADay   = 24
// 	hoursInAMonth = 730
// 	hoursInAYear  = 8760
// )

type Item struct {
	Title  string `json:"title"`
	Author string `json:"by"`
	// Score  int    `json:"score"`
	// ID     int    `json:"id"`
	// time   time.Time
	// absolutePostTime  time.Time
	// timeSincePosting time.Duration
	// URL string `json:"url"`
}

type Items []Item

// var HItem Item
// var HItems Items

func main() {
	fmt.Println("Hello hn-cli user")

	dataToMarshall := Item{"Alice in Wonderland", "Lewis Carroll"}

	dataMarshalled, mErr := Marshall(dataToMarshall)
	if mErr != nil {
		panic(mErr)
	}
	// debug
	fmt.Println(dataMarshalled)

	// dataToUnmarshall := Item{Title: "Alice in Wonderland", Author: "Lewis Carroll"}
	dataUnmarshalled, uErr := Unmarshal(dataMarshalled)
	if uErr != nil {
		panic(uErr)
	}
	// debug
	fmt.Println(dataUnmarshalled)
}

func Marshall(input Item) ([]byte, error) {
	dataMarshalled, err := json.Marshal(input)
	if err != nil {
		return dataMarshalled, err
	}

	return dataMarshalled, nil
}

func Unmarshal(input []byte) (Item, error) {
	var dataUnmarshalled Item

	err := json.Unmarshal(input, &dataUnmarshalled)
	if err != nil {
		return dataUnmarshalled, err
	}

	return dataUnmarshalled, nil
}

//	func (t Item) addHoursSincePosting() time.Duration {
//		t.hoursSincePosting = time.Since(t.absolutePostTime)
//		return t.hoursSincePosting
//	}
// func (t Item) relativeTime() string {
// 	elapsedHours := t.timeSincePosting.Hours()

// 	if elapsedHours < 1 {
// 		return "<1h"
// 	}

// 	if elapsedHours < hoursInADay {
// 		return fmt.Sprint(math.Round(elapsedHours), "h")
// 	}

// 	if elapsedHours < hoursInAMonth {
// 		return fmt.Sprint(math.Round(elapsedHours/hoursInADay), "d")
// 	}

// 	if elapsedHours < hoursInAYear {
// 		return fmt.Sprint(math.Round(elapsedHours/hoursInAMonth), "m")
// 	}

// 	return fmt.Sprint(math.Round(elapsedHours/hoursInAYear), "y")
// }
