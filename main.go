package main

import (
	"fmt"
	"log"

	"github.com/BenchR267/lbd/lexer"
)

func main() {
	c, err := lexer.StreamFromFile("example/main.lbd")
	if err != nil {
		log.Fatal(err)
	}

	l := lexer.NewLexer(c)
	l.Start()
	for t := range l.NextToken {
		fmt.Println(t)
	}
}
