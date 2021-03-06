package lexer

import (
	"testing"

	"github.com/BenchR267/lbd/lexer/token"
)

func TestTokenizer(t *testing.T) {
	testCases := []struct {
		input    string
		expected token.Type
	}{
		{"abc", token.Identifier},
		{"a", token.Identifier},

		{"5", token.Integer},

		{"(", token.ParenthesisOpen},
		{")", token.ParenthesisClose},
		{"{", token.CurlyBracketOpen},
		{"}", token.CurlyBracketClose},
		{"[", token.SquareBracketOpen},
		{"]", token.SquareBracketClose},

		{"->", token.Arrow},
		{",", token.Comma},

		{"=", token.Assign},
		{"+", token.Plus},
		{"-", token.Minus},
		{"/", token.Slash},
		{"%", token.Percent},

		{"==", token.Equal},
		{"!=", token.NotEqual},
		{">", token.Greater},
		{"<", token.Less},
		{">=", token.GreaterEqual},
		{"<=", token.LessEqual},

		{"#", token.Illegal},
	}

	for _, c := range testCases {
		tkn := tokenizer{
			content: []rune{},
		}
		for i := 0; i < len(c.input); i++ {
			got := tkn.append(rune(c.input[i]), token.Position{})
			if got != nil {
				t.Errorf("Expected token to be nil, but got %#v instead. input: %s", got, c.input)
			}
		}

		got := tkn.append(' ', token.Position{})
		if got.Type != c.expected {
			t.Errorf("Expected to get tokentype %#v, but got %#v instead.", c.expected.String(), got.Type.String())
		}
	}
}
