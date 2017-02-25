package main

import (
	"fmt"

	"github.com/BenchR267/lbd/lexer"
)

func main() {
	c := lexer.StreamFromString(`
method = (a int, b int) -> int {
	return a + b
}
`)

	l := lexer.NewLexer(c)
	l.Start()
	for t := range l.NextToken {
		fmt.Println(t)
	}
}
