package lexer

import (
	"unicode"

	"github.com/BenchR267/lbd/lexer/token"
)

type tokenizer struct {
	content []rune
}

func (t *tokenizer) append(by rune, pos token.Position) *token.Token {
	content := t.content
	defer func() {
		t.content = append(t.content, by)
	}()

	if !belongsTogether(content, by) {
		return t.token(pos)
	}
	return nil
}

func (t *tokenizer) token(currentPosition token.Position) *token.Token {
	content := string(t.content)
	if len(content) == 0 {
		return nil
	}

	t.content = []rune{}
	p := currentPosition
	p.Column -= len(content)
	p.Len = len(content)
	return &token.Token{
		Pos:  p,
		Raw:  content,
		Type: token.FromRaw(content),
	}
}

func belongsTogether(current []rune, next rune) bool {
	if len(current) == 0 {
		return true
	}

	if len(current) == 1 {
		switch current[0] {
		case '+':
			return next == '+'
		case '!', '=':
			return next == '='
		case '<':
			return next == '='
		case '>':
			return next == '='
		case '-':
			return next == '>'
		}
		return unicode.IsLetter(current[0]) && unicode.IsLetter(next)
	} else if len(current) == 2 {
		s := string(current)
		t := token.FromRaw(s)

		if t != token.Identifier && t != token.Integer && t != token.Illegal {
			return t == token.Illegal
		}
		return unicode.IsLetter(rune(next))
	} else {
		return unicode.IsLetter(rune(next)) && token.IsLetter(string(current))
	}
}
