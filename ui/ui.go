package ui

import "fmt"

func UI() (string, error) {
	var input string

	fmt.Printf("Type `start' to display posts, 'help' for list of commands\n>")

	_, err := fmt.Scan(&input)

	return input, err
}
