package lexer

import "github.com/BenchR267/lbd/lexer/token"

type Lexer struct {
	input <-chan byte

	NextToken chan token.Token

	currentPosition token.Position

	buffer Tokenizer
}

func NewLexer(inputStream <-chan byte) *Lexer {
	l := &Lexer{
		input:     inputStream,
		NextToken: make(chan token.Token),
		currentPosition: token.Position{
			Column: 0,
			Line:   0,
		},
		buffer: Tokenizer{
			content: []byte{},
		},
	}
	return l
}

func (l Lexer) Start() {
	go func() {
		for b := range l.input {

			if !isWhitespace(b) {
				t := l.buffer.append(b, l.currentPosition)
				if t != nil {
					l.NextToken <- *t
				}
			} else {
				t := l.buffer.token(l.currentPosition)
				if t != nil {
					l.NextToken <- *t
				}
			}

			if b == '\n' {
				l.currentPosition.Column = 0
				l.currentPosition.Line++
			} else {
				l.currentPosition.Column++
			}

		}
		t := l.buffer.token(l.currentPosition)
		if t != nil {
			l.NextToken <- *t
		}
		close(l.NextToken)
	}()
}
