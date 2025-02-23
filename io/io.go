package io

import (
	"golang.org/x/term"
)

func TermSize() (int, int, error) {
	tWidth, tHeight, err := term.GetSize(0)
	return tWidth, tHeight, err
}
