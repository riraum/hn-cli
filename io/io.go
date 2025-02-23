package io

import (
	"golang.org/x/term"
)

func GetSize() (int, int, error) {
	tWidth, tHeight, tErr := term.GetSize(0)
	if tErr != nil {
		return tWidth, tHeight, tErr
	}

	return tWidth, tHeight, nil
}
