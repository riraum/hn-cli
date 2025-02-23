package ui

import "fmt"

func UI() (string, error) {
	var input string

	fmt.Printf("Type 'help' for list of commands\n>")

	_, iErr := fmt.Scan(&input)
	if iErr != nil {
		return input, iErr
	}

	return input, nil
}
