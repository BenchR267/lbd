package token

import "unicode"

// Type defines the type of a token
type Type int

// Position defines the position in the source code
type Position struct {
	Line   int
	Column int
}

// Token defines one token in the source code
type Token struct {
	Pos Position

	Type Type
	Raw  string
}

// All possible token types
const (
	Identifier Type = iota

	Integer
	Float

	Parenthesis
	CurlyBracket
	SquareBracket

	Assign
	Plus
	Minus
	Slash
	Percent

	Equal
	NotEqual
	Greater
	Less

	Illegal
)

func FromRaw(raw string) Type {
	switch raw {
	case "(":
		fallthrough
	case ")":
		return Parenthesis
	case "{":
		fallthrough
	case "}":
		return CurlyBracket
	case "[":
		fallthrough
	case "]":
		return SquareBracket
	case "=":
		return Assign
	case "+":
		return Plus
	case "-":
		return Minus
	case "/":
		return Slash
	case "%":
		return Percent
	case "==":
		return Equal
	case "!=":
		return NotEqual
	case ">":
		return Greater
	case "<":
		return Less
	}
	if IsLetter(raw) {
		return Identifier
	} else if isInteger(raw) {
		return Integer
	}
	return Illegal
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isInteger(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
