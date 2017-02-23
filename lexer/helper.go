package lexer

import "unicode"

func StreamFromString(input string) <-chan byte {
	c := make(chan byte)

	go func() {
		for _, b := range []byte(input) {
			c <- b
		}
		close(c)
	}()

	return c
}

func isWhitespace(b byte) bool {
	return b == '\n' || b == ' '
}

func isLetter(b byte) bool {
	return unicode.IsLetter(rune(b))
}
