package ui

import "fmt"

func UI() ([]string, error) {
	var input string

	var inputNum string

	fmt.Printf("Type `start' to display posts, 'help' for list of commands\n>")

	_, err := fmt.Scanf("%s %s", &input, &inputNum)
	fmt.Println("debug:", input, inputNum)

	if err != nil {
		if inputNum == "" {
			return []string{input}, err
		}

		return []string{input, inputNum}, err
	}

	if inputNum == "" {
		return []string{input}, nil
	}

	return []string{input, inputNum}, nil
}
