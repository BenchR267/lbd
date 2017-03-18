package lexer

import (
	"errors"

	"github.com/BenchR267/lbd/lexer/token"
)

var (
	// ErrNotFinished is returned on Start() when Lexer was started before and has not finished yet
	ErrNotFinished = errors.New("lexer still not finished lexing")

	// ErrInputStreamNil is returned on Start() if the given inputStream is nil
	ErrInputStreamNil = errors.New("input stream should not be nil")
)

// Lexer represents an instance to get a lexical representation of the source code.
//
// It works in it's own go routine, so after creation with NewLexer get the tokens via
// the NextToken field.
type Lexer struct {
	NextToken chan token.Token

	input  <-chan rune
	curPos token.Position
	buffer tokenizer
}

// NewLexer creates a new instance of Lexer, ready to be started.
func NewLexer() *Lexer {
	l := &Lexer{
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

// Start will read from the inputStream, forwarding tokens via NextToken.
// Start runs in its own go routine and will get a zombie if NextToken is not read!
func (l *Lexer) Start(inputStream <-chan rune) error {
	if l.input != nil {
		return ErrNotFinished
	}
	if inputStream == nil {
		return ErrInputStreamNil
	}
	l.input = inputStream
	l.NextToken = make(chan token.Token)
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
		l.input = nil
		close(l.NextToken)
	}()
	return nil
}
