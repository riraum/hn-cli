package ui

import "fmt"

func UI() (string, int, error) {
	var input string

	var inputInt int

	fmt.Printf("Type 'help' for list of commands\n>")

	_, err := fmt.Scanf("%s %d", &input, &inputInt)
	fmt.Println("debug:", input, inputInt)

	return input, inputInt, err
}
