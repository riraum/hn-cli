package io

import (
	"fmt"

	"golang.org/x/term"
)

func TermSize() (int, int, error) {
	var int1 int

	var int2 int
	if int1, _, err := term.GetSize(0); err != nil {
		return int1, int2, fmt.Errorf("Failed to get terminal size: %w", err)
	}

	return int1, int2, nil
}
