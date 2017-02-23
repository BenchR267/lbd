package lexer

import (
	"unicode"

	"github.com/BenchR267/lbd/lexer/token"
)

type Tokenizer struct {
	content []byte
}

func (b *Tokenizer) append(by byte, pos token.Position) *token.Token {

	content := b.content

	defer func() {
		b.content = append(b.content, by)
	}()

	if !belongsTogether(content, by) {
		return b.token(pos)
	}
	return nil
}

func (b *Tokenizer) token(currentPosition token.Position) *token.Token {
	content := string(b.content)
	b.content = []byte{}
	return &token.Token{
		Pos:  currentPosition,
		Raw:  content,
		Type: token.FromRaw(content),
	}
}

func belongsTogether(current []byte, next byte) bool {
	if len(current) == 0 {
		return true
	}

	if len(current) == 1 {
		switch current[0] {
		case '+':
			return next == '+'
		case '!':
			fallthrough
		case '=':
			return next == '='
		}
		return isLetter(current[0]) && isLetter(next)
	} else if len(current) == 2 {
		s := string(current)
		t := token.FromRaw(s)

		if t != token.Identifier && t != token.Integer && t != token.Float && t != token.Illegal {
			return t == token.Illegal
		} else {
			return unicode.IsLetter(rune(next))
		}
	} else {
		return unicode.IsLetter(rune(next)) && token.IsLetter(string(current))
	}

	return false
}
