package ui

import "fmt"

func UI() []string {
	var input string

	var inputNum string

	fmt.Printf("Type `start' to display posts, 'help' for list of commands\n>")

	_, err := fmt.Scanf("%s %s", &input, &inputNum)
	fmt.Println("debug:", input, inputNum)

	if inputNum == "" {
		return []string{input}
	}

	if err != nil {
		panic(err)
	}

	rv := []string{input, inputNum}
	fmt.Println(rv)

	return rv
}
