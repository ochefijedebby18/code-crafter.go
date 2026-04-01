package main

import (
	"strings"
)

func cases(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])

			words = append(words[:i], words[i+1:]...)
			i--
		}
		if words[i] == "(low)" && i > 0 {
			words[i-1] = strings.ToLower(words[i-1])

			words = append(words[:i], words[i+1:]...)
			i--
		}
		if words[i] == "(cap)" && i > 0 {
			words[i-1] = strings.Title(words[i-1])

			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}

