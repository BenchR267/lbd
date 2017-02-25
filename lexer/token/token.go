package token

import "unicode"

// Type defines the type of a token.
type Type int

// Position defines the position in the source code.
type Position struct {
	Line   int
	Column int
	Len    int
}

// Token defines one token in the source code.
type Token struct {
	Pos  Position
	Type Type
	Raw  string
}

// All possible token types.
const (
	Identifier Type = iota

	Keyword
	BuildInType

	Integer

	ParenthesisOpen
	ParenthesisClose

	CurlyBracketOpen
	CurlyBracketClose

	SquareBracketOpen
	SquareBracketClose

	Arrow
	Comma

	Assign
	Plus
	Minus
	Slash
	Multiply
	Percent

	Equal
	NotEqual
	Greater
	Less
	GreaterEqual
	LessEqual

	Illegal
)

const (
	// Return is used to return from a function.
	Return = "return"
)

const (
	// Int describes an integer variable type (should maybe moved out of token to parser).
	Int = "int"
)

var keywords = map[string]bool{
	Return: true,
}

var types = map[string]bool{
	Int: true,
}

// FromRaw returns the token.Type for the given raw value.
func FromRaw(raw string) Type {
	switch raw {
	case "(":
		return ParenthesisOpen
	case ")":
		return ParenthesisClose
	case "{":
		return CurlyBracketOpen
	case "}":
		return CurlyBracketClose
	case "[":
		return SquareBracketOpen
	case "]":
		return SquareBracketClose
	case "->":
		return Arrow
	case ",":
		return Comma
	case "=":
		return Assign
	case "+":
		return Plus
	case "-":
		return Minus
	case "/":
		return Slash
	case "*":
		return Multiply
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
	case ">=":
		return GreaterEqual
	case "<=":
		return LessEqual
	}
	if IsLetter(raw) {
		if IsKeyword(raw) {
			return Keyword
		} else if IsBuildInType(raw) {
			return BuildInType
		}
		return Identifier
	}
	if IsInteger(raw) {
		return Integer
	}
	return Illegal
}

// IsLetter returns true if the given string contains only letters
func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// IsInteger returns true if the given string contains only digits
func IsInteger(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// IsKeyword returns true if the given string is a reserved keyword of the language
func IsKeyword(s string) bool {
	_, ok := keywords[s]
	return ok
}

// IsBuildInType returns true if the given string is a reserved built in type
func IsBuildInType(s string) bool {
	_, ok := types[s]
	return ok
}
