package main

import "github.com/BenchR267/lbd/lexer"
import "fmt"

func main() {

	c := lexer.StreamFromString(`abc == b
c + d
b++
f=5`)

	l := lexer.NewLexer(c)

	l.Start()

	for t := range l.NextToken {
		fmt.Println(t)
	}
}
