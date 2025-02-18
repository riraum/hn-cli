package main

import (
	"fmt"

	"github.com/riraum/hn-cli/item"
)

func main() {
	fmt.Println("Hello hn-cli user")

	dataToMarshall := item.Item{Title: "Alice in Wonderland", Author: "Lewis Carroll"}

	dataMarshalled, mErr := item.Marshall(dataToMarshall)
	if mErr != nil {
		panic(mErr)
	}
	// debug
	fmt.Println(dataMarshalled)

	dataUnmarshalled, uErr := item.Unmarshal(dataMarshalled)
	if uErr != nil {
		panic(uErr)
	}
	// debug
	fmt.Println(dataUnmarshalled)
}
