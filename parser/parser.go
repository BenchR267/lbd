package parser

import (
	"fmt"

	"github.com/BenchR267/lbd/lexer"
	"github.com/BenchR267/lbd/lexer/token"
)

// Parser creates an AST from a given stream of tokens
type Parser struct {
	lexer *lexer.Lexer
}

// ParseError wraps information about an error that occured during parsing
type ParseError struct {
	Expected token.Token
	Got      token.Token
}

func (e ParseError) Error() string {
	return fmt.Sprintf("Expected %s, but got %s.", e.Expected.Raw, e.Got.Raw)
}
