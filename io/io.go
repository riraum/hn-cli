package io

import (
	"fmt"

	"golang.org/x/term"
)

func TermSize() (int, error) {
	tWidth, _, err := term.GetSize(0)
	if err != nil {
		return tWidth, fmt.Errorf("Failed to get terminal width %w", err)
	}
	return tWidth, nil
}
