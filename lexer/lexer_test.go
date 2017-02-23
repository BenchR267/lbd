package lexer

import (
	"testing"

	"github.com/BenchR267/lbd/lexer/token"
)

func TestStreamFromString(t *testing.T) {
	count := 0

	const input = "abc"

	stream := StreamFromString(input)

	for b := range stream {
		if b != input[count] {
			t.Errorf("Expected to receive %c over channel at index %d from input %s", b, count, input)
		}
		count++
	}

	if count != len(input) {
		t.Errorf("Expected count (%d) to be equal to %d.", count, len(input))
	}
}

func TestIsWhitespace(t *testing.T) {
	testCases := []struct {
		input    byte
		expected bool
	}{
		{'a', false},
		{'b', false},
		{'\n', true},
		{' ', true},
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
		input    byte
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

func TestTokenizer(t *testing.T) {
	testCases := []struct {
		input    string
		expected token.Type
	}{
		{"abc", token.Identifier},
	}

	for _, c := range testCases {
		tokenizer := Tokenizer{
			content: []byte{},
		}
		for i := 0; i < len(c.input); i++ {
			got := tokenizer.append(c.input[i], token.Position{})
			if got != nil {
				t.Errorf("Expected token to be nil, but got %#v instead. input: %s", got, c.input)
			}
		}

		got := tokenizer.append(' ', token.Position{})
		if got.Type != c.expected {
			t.Errorf("Expected to get tokentype %#v, but got %#v instead.", c.expected.String(), got.Type.String())
		}
	}
}
