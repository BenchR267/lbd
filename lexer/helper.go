package lexer

import (
	"bufio"
	"os"
	"unicode"
)

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

func StreamFromFile(filename string) (stream <-chan rune, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}

	fileStream := make(chan rune)

	go func() {
		reader := bufio.NewReader(file)
		for {
			r, _, err := reader.ReadRune()
			if err != nil {
				close(fileStream)
				break
			}
			fileStream <- r
		}
	}()

	stream = fileStream

	return
}

func CombineStreams(streams ...<-chan rune) <-chan rune {
	newChan := make(chan rune)
	go func() {
		for _, s := range streams {
			for r := range s {
				newChan <- r
			}
		}
		close(newChan)
	}()
	return newChan
}

func isWhitespace(b rune) bool {
	return b == '\n' || b == ' ' || b == '\t'
}

func isLetter(b rune) bool {
	return unicode.IsLetter(b)
}
