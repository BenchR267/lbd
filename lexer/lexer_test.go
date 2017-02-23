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
		{"a", token.Identifier},

		{"5", token.Integer},

		{"(", token.Parenthesis},
		{"{", token.CurlyBracket},
		{"[", token.SquareBracket},

		{"=", token.Assign},
		{"+", token.Plus},
		{"-", token.Minus},
		{"/", token.Slash},
		{"%", token.Percent},

		{"==", token.Equal},
		{"!=", token.NotEqual},
		{">", token.Greater},
		{"<", token.Less},

		{"#", token.Illegal},
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

func TestStart_OneLine(t *testing.T) {
	stream := StreamFromString("abc = dfe + 3")
	l := NewLexer(stream)
	l.Start()

	expectedValues := []token.Token{
		token.Token{
			Pos:  token.Position{Line: 0, Column: 0, Len: 3},
			Type: token.Identifier,
			Raw:  "abc",
		},
		token.Token{
			Pos:  token.Position{Line: 0, Column: 4, Len: 1},
			Type: token.Assign,
			Raw:  "=",
		},
		token.Token{
			Pos:  token.Position{Line: 0, Column: 6, Len: 3},
			Type: token.Identifier,
			Raw:  "dfe",
		},
		token.Token{
			Pos:  token.Position{Line: 0, Column: 10, Len: 1},
			Type: token.Plus,
			Raw:  "+",
		},
		token.Token{
			Pos:  token.Position{Line: 0, Column: 12, Len: 1},
			Type: token.Integer,
			Raw:  "3",
		},
	}

	i := 0
	for token := range l.NextToken {
		expected := expectedValues[i]

		if expected != token {
			t.Errorf("Expected: %#v, got: %#v.", expected, token)
		}

		i++
	}
}

func TestStart_MultipleLines(t *testing.T) {
	stream := StreamFromString(`a = 5
b = 4
c = a + b`)
	l := NewLexer(stream)
	l.Start()

	expectedValues := []token.Token{
		token.Token{
			Pos:  token.Position{Line: 0, Column: 0, Len: 1},
			Type: token.Identifier,
			Raw:  "a",
		},
		token.Token{
			Pos:  token.Position{Line: 0, Column: 2, Len: 1},
			Type: token.Assign,
			Raw:  "=",
		},
		token.Token{
			Pos:  token.Position{Line: 0, Column: 4, Len: 1},
			Type: token.Integer,
			Raw:  "5",
		},
		token.Token{
			Pos:  token.Position{Line: 1, Column: 0, Len: 1},
			Type: token.Identifier,
			Raw:  "b",
		},
		token.Token{
			Pos:  token.Position{Line: 1, Column: 2, Len: 1},
			Type: token.Assign,
			Raw:  "=",
		},
		token.Token{
			Pos:  token.Position{Line: 1, Column: 4, Len: 1},
			Type: token.Integer,
			Raw:  "4",
		},
		token.Token{
			Pos:  token.Position{Line: 2, Column: 0, Len: 1},
			Type: token.Identifier,
			Raw:  "c",
		},
		token.Token{
			Pos:  token.Position{Line: 2, Column: 2, Len: 1},
			Type: token.Assign,
			Raw:  "=",
		},
		token.Token{
			Pos:  token.Position{Line: 2, Column: 4, Len: 1},
			Type: token.Identifier,
			Raw:  "a",
		},
		token.Token{
			Pos:  token.Position{Line: 2, Column: 6, Len: 1},
			Type: token.Plus,
			Raw:  "+",
		},
		token.Token{
			Pos:  token.Position{Line: 2, Column: 8, Len: 1},
			Type: token.Identifier,
			Raw:  "b",
		},
	}

	i := 0
	for token := range l.NextToken {
		expected := expectedValues[i]

		if expected != token {
			t.Errorf("Expected: %#v, got: %#v.", expected, token)
		}

		i++
	}
}
