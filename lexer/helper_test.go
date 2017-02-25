package lexer

import "testing"
import "unicode/utf8"
import "os"

func TestStreamFromString(t *testing.T) {
	count := 0

	const input = "abc"

	stream := StreamFromString(input)

	for r := range stream {
		if r != rune(input[count]) {
			t.Errorf("Expected to receive %c over channel at index %d from input %s", r, count, input)
		}
		count++
	}

	if count != len(input) {
		t.Errorf("Expected count (%d) to be equal to %d.", count, len(input))
	}
}

func TestStreamFromFile(t *testing.T) {
	_, err := StreamFromFile("aNonExistingFile.lbd")
	if err == nil {
		t.Error("Expected to get error when opening non existing file, but got nil instead.")
	}

	const fileName = "../src/main.lbd"

	size, err := fileSize(fileName)
	if err != nil {
		t.Errorf("Expected to get no error when opening existing file, but got %s instead.", err.Error())
	}

	s, err := StreamFromFile(fileName)
	if err != nil {
		t.Errorf("Expected to get no error when opening existing file, but got %s instead.", err.Error())
	}
	if s == nil {
		t.Error("Expected stream of existing file to be non nil.")
	}

	i := int64(0)
	for r := range s {
		i += int64(utf8.RuneLen(r))
	}

	if i != size {
		t.Errorf("Expected to get filesize of 85 instead of %d", i)
	}
}

func fileSize(fileName string) (size int64, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	fi, err := f.Stat()
	if err != nil {
		return
	}
	size = fi.Size()
	return
}

func TestCombineStreams(t *testing.T) {
	s1 := StreamFromString("a = 4")
	s2 := StreamFromString("b = 5")
	if s1 == nil || s2 == nil {
		t.Error("Expected both streams to be non nil")
	}

	combined := CombineStreams(s1, s2)

	runes := []rune{'a', ' ', '=', ' ', '4', 'b', ' ', '=', ' ', '5'}

	i := 0
	for r := range combined {
		if r != runes[i] {
			t.Errorf("Expected to get %c at index %d, but got %c instead.", runes[i], i, r)
		}
		i++
	}

	if i != len(runes) {
		t.Errorf("Expected to get %d runes, but got %s instead.", len(runes), i)
	}
}

func TestIsWhitespace(t *testing.T) {
	testCases := []struct {
		input    rune
		expected bool
	}{
		{'a', false},
		{'b', false},
		{'\n', true},
		{' ', true},
		{'\t', true},
	}

	for _, c := range testCases {
		got := isWhitespace(c.input)
		if got != c.expected {
			t.Errorf("Whitespace test fail. Input: %c - Expected: %#v - Got: %#v", c.input, c.expected, got)
		}
	}
}

func TestIsLetter(t *testing.T) {
	testCases := []struct {
		input    rune
		expected bool
	}{
		{'a', true},
		{'b', true},
		{'\n', false},
		{' ', false},
		{'3', false},
		{'$', false},
	}

	for _, c := range testCases {
		got := isLetter(c.input)
		if got != c.expected {
			t.Errorf("Letter test fail. Input: %c - Expected: %#v - Got: %#v", c.input, c.expected, got)
		}
	}
}
