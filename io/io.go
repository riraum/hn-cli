package io

import (
	"golang.org/x/term"
)

func TermSize() (int, int, error) {
	tWidth, tHeight, err := term.GetSize(0)
	if err != nil {
		return tWidth, tHeight, err
	}

	return tWidth, tHeight, nil
}
