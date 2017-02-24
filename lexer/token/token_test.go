package token

import "testing"

func TestFromRaw(t *testing.T) {
	tests := []struct {
		input    string
		expected Type
	}{
		{"(", Parenthesis},
		{")", Parenthesis},
		{"{", CurlyBracket},
		{"}", CurlyBracket},
		{"[", SquareBracket},
		{"]", SquareBracket},
		{"->", Arrow},
		{",", Comma},
		{"=", Assign},
		{"+", Plus},
		{"-", Minus},
		{"/", Slash},
		{"%", Percent},
		{"==", Equal},
		{"!=", NotEqual},
		{">", Greater},
		{"<", Less},
		{">=", GreaterEqual},
		{"<=", LessEqual},
		{"$", Illegal},
		{"aVariable", Identifier},
		{"125", Integer},
	}

	for _, test := range tests {
		got := FromRaw(test.input)
		if got != test.expected {
			t.Errorf("Expected %s, got %s.", test.expected.String(), got.String())
		}
	}
}

func TestTypeString(t *testing.T) {
	tests := []struct {
		input    Type
		expected string
	}{
		{Identifier, "Identifier"},
		{Integer, "Integer"},
		{Parenthesis, "Parenthesis"},
		{CurlyBracket, "CurlyBracket"},
		{SquareBracket, "SquareBracket"},
		{Arrow, "Arrow"},
		{Assign, "Assign"},
		{Plus, "Plus"},
		{Minus, "Minus"},
		{Slash, "Slash"},
		{Percent, "Percent"},
		{Equal, "Equal"},
		{NotEqual, "NotEqual"},
		{Greater, "Greater"},
		{Less, "Less"},
		{GreaterEqual, "GreaterEqual"},
		{LessEqual, "LessEqual"},
		{Illegal, "Illegal"},
		{Type(100), "Type(100)"},
	}

	for _, test := range tests {
		got := test.input.String()
		if got != test.expected {
			t.Errorf("Expected %s, got %s.", test.expected, got)
		}
	}
}
