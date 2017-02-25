package lexer

import (
	"bufio"
	"os"
)

// StreamFromString returns a read only rune channel representing the given string.
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

// StreamFromFile returns a read only rune channel representing the given file content.
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
		file.Close()
	}()

	stream = fileStream

	return
}

// CombineStreams creates a sequence of the given streams, processing all successively.
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

// isWhitespace defines if a given rune should be handled as whitespace.
func isWhitespace(b rune) bool {
	return b == '\n' || b == ' ' || b == '\t'
}
