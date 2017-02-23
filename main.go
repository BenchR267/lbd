package main

import "github.com/BenchR267/lbd/lexer"
import "fmt"

func main() {

	c := lexer.StreamFromString(`
abc = 1
b = 4
c = abc + b
`)

	l := lexer.NewLexer(c)

	l.Start()

	for t := range l.NextToken {
		fmt.Println(t)
	}
}
