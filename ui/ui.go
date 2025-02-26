package ui

import "fmt"

func UI() (string, int, bool, error) {
	var input string

	var inputInt int

	fmt.Printf("Type 'help' for list of commands\n>")

	_, err := fmt.Scanf("%s %d", &input, &inputInt)
	fmt.Println("debug:", input, inputInt)
	if inputInt == 0 {
		return input, inputInt, true, err
	}

	return input, inputInt, false, err
}
