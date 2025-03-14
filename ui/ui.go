package ui

import (
	"fmt"
)

func UI() ([]string, error) {
	var input string

	var inputNum string

	fmt.Printf("Type `start' to display posts, 'help' for list of commands\n>")

	_, err := fmt.Scanf("%s %s", &input, &inputNum)

	if input == "" {
		return []string{""}, fmt.Errorf("Empty input %w", err)
	}

	if err != nil {
		if inputNum == "" {
			return []string{input}, fmt.Errorf("Failed to get string input (no index entered) %w", err)
		}

		return []string{input, inputNum}, fmt.Errorf("Failed to get string and index input %w", err)
	}

	if inputNum == "" {
		fmt.Println("input:", input)
		return []string{input}, nil
	}

	return []string{input, inputNum}, nil
}
