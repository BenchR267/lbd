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

// New creates a new instance of Lexer, ready to be started.
func New() *Lexer {
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
func (lex *Lexer) Start(inputStream <-chan rune) error {
	if lex.input != nil {
		return ErrNotFinished
	}
	if inputStream == nil {
		return ErrInputStreamNil
	}
	lex.input = inputStream
	lex.NextToken = make(chan token.Token)
	go func() {
		for b := range lex.input {

			var t *token.Token

			if !isWhitespace(b) {
				t = lex.buffer.append(b, lex.curPos)
			} else {
				t = lex.buffer.token(lex.curPos)
			}

			if t != nil {
				lex.NextToken <- *t
			}

			if b == '\n' {
				lex.curPos.Column = 0
				lex.curPos.Line++
			} else {
				lex.curPos.Column++
			}

		}

		t := lex.buffer.token(lex.curPos)
		if t != nil {
			lex.NextToken <- *t
		}

		lex.input = nil
		close(lex.NextToken)
	}()
	return nil
}
