package lexer

import (
	"testing"

	"github.com/BenchR267/lbd/lexer/token"
)

func TestIgnoreWhitespace(t *testing.T) {
	stream := StreamFromString("a=    5")
	l := NewLexer(stream)
	l.Start()

	expectedValues := []token.Token{
		token.Token{
			Pos:  token.Position{Line: 0, Column: 0, Len: 1},
			Type: token.Identifier,
			Raw:  "a",
		},
		token.Token{
			Pos:  token.Position{Line: 0, Column: 1, Len: 1},
			Type: token.Assign,
			Raw:  "=",
		},
		token.Token{
			Pos:  token.Position{Line: 0, Column: 6, Len: 1},
			Type: token.Integer,
			Raw:  "5",
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

func TestConditions(t *testing.T) {
	stream := StreamFromString("a<=b")
	l := NewLexer(stream)
	l.Start()

	expectedValues := []token.Token{
		token.Token{
			Pos:  token.Position{Line: 0, Column: 0, Len: 1},
			Type: token.Identifier,
			Raw:  "a",
		},
		token.Token{
			Pos:  token.Position{Line: 0, Column: 1, Len: 2},
			Type: token.LessEqual,
			Raw:  "<=",
		},
		token.Token{
			Pos:  token.Position{Line: 0, Column: 3, Len: 1},
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

func TestFunction(t *testing.T) {
	stream := StreamFromString(`
a = (a int, b int) -> int {
	return a + b
}
	`)

	l := NewLexer(stream)
	l.Start()

	expectedValues := []struct {
		Type token.Type
		Raw  string
	}{
		{token.Identifier, "a"},
		{token.Assign, "="},
		{token.Parenthesis, "("},
		{token.Identifier, "a"},
		{token.BuildInType, "int"},
		{token.Comma, ","},
		{token.Identifier, "b"},
		{token.BuildInType, "int"},
		{token.Parenthesis, ")"},
		{token.Arrow, "->"},
		{token.BuildInType, "int"},
		{token.CurlyBracketOpen, "{"},
		{token.Keyword, "return"},
		{token.Identifier, "a"},
		{token.Plus, "+"},
		{token.Identifier, "b"},
		{token.CurlyBracketClose, "}"},
	}

	i := 0
	for token := range l.NextToken {
		expected := expectedValues[i]

		if expected.Raw != token.Raw {
			t.Errorf("Expected Raw: %#v, got: %#v.", expected.Raw, token.Raw)
		}

		if expected.Type != token.Type {
			t.Errorf("Expected Type: %#v, got: %#v.", expected.Type, token.Type)
		}

		i++
	}
}
