package lexer

import "github.com/BenchR267/lbd/lexer/token"

type Lexer struct {
	NextToken chan token.Token

	input  <-chan rune
	curPos token.Position
	buffer tokenizer
}

func NewLexer(inputStream <-chan rune) *Lexer {
	l := &Lexer{
		NextToken: make(chan token.Token),
		input:     inputStream,
		curPos: token.Position{
			Column: 0,
			Line:   0,
		},
		buffer: tokenizer{
			content: []rune{},
		},
	}
	return l
}

func (l *Lexer) Start() {
	go func() {
		for b := range l.input {
			if !isWhitespace(b) {
				t := l.buffer.append(b, l.curPos)
				if t != nil {
					l.NextToken <- *t
				}
			} else {
				t := l.buffer.token(l.curPos)
				if t != nil {
					l.NextToken <- *t
				}
			}

			if b == '\n' {
				l.curPos.Column = 0
				l.curPos.Line++
			} else {
				l.curPos.Column++
			}

		}
		t := l.buffer.token(l.curPos)
		if t != nil {
			l.NextToken <- *t
		}
		close(l.NextToken)
	}()
}
