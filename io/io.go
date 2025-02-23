package io

import (
	"golang.org/x/term"
)

func TermSize() (int, int, error) {
	return term.GetSize(0)
}
