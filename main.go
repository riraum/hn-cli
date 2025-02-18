package main

import (
	"encoding/json"
	"fmt"
)

const (
	hoursInADay   = 24
	hoursInAMonth = 730
	hoursInAYear  = 8760
)

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
	fmt.Println("Hello hn-cli")

	dataToMarshall := Item{"Alice in Wonderland", "Lewis Carroll"}
	// dataMarshalled, mErr := json.Marshal(dataToMarshall)
	// if mErr != nil {
	// 	panic(mErr)
	// }
	dataMarshalled := Marshall(dataToMarshall)
	// debug
	// fmt.Println(dataMarshalled)
	fmt.Println("Marshalled data:", string(dataMarshalled))

	// dataToUnmarshall := Item{Title: "Alice in Wonderland", Author: "Lewis Carroll"}
	var dataUnmarshalled Item
	// uErr := json.Unmarshal(dataMarshalled, &dataUnmarshalled)
	// if uErr != nil {
	// 	panic(uErr)
	// }
	dataUnmarshalled = Unmarshal(dataMarshalled)
	fmt.Println("Unmarshalled data:", dataUnmarshalled)

}

func Marshall(input Item) []byte {
	dataMarshalled, mErr := json.Marshal(input)
	if mErr != nil {
		panic(mErr)
	}
	return dataMarshalled
}

func Unmarshal(input []byte) Item {
	// stringBytes := []byte(input)
	var dataUnmarshalled Item
	uErr := json.Unmarshal(input, &dataUnmarshalled)
	if uErr != nil {
		panic(uErr)
	}
	return dataUnmarshalled
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
