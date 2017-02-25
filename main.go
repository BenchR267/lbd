package main

import "github.com/BenchR267/lbd/lexer"
import "fmt"
import "log"

func main() {

	c, err := lexer.StreamFromFile("src/main.lbd")
	if err != nil {
		log.Fatal(err)
	}

	l := lexer.NewLexer(c)

	l.Start()

	for t := range l.NextToken {
		fmt.Println(t)
	}
}
