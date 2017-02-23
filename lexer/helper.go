package lexer

import "unicode"

func StreamFromString(input string) <-chan rune {
	c := make(chan rune)

	go func() {
		for _, b := range input {
			c <- b
		}
		close(c)
	}()

	return c
}

func isWhitespace(b rune) bool {
	return b == '\n' || b == ' '
}

func isLetter(b rune) bool {
	return unicode.IsLetter(b)
}
