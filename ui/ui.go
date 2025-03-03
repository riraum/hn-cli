package ui

import "fmt"

func UI() ([]string, error) {
	var input string

	var inputNum string

	fmt.Printf("Type `start' to display posts, 'help' for list of commands\n>")

	_, err := fmt.Scanf("%s %s", &input, &inputNum)
	if err != nil {
		if inputNum == "" {
			return nil, fmt.Errorf("Failed to get `%s` string input (no index entered) %w", []string{input}, err)
		}

		return nil, fmt.Errorf("Failed to get `%s` string and index input %w", []string{input, inputNum}, err)
	}

	if inputNum == "" {
		return []string{input}, nil
	}

	return []string{input, inputNum}, nil
}
