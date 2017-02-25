package repl

import (
	"fmt"
	"os"

	"bufio"

	"github.com/BenchR267/lbd/lexer"
)

// Start is starting the interactive REPL (currently just printing out tokens)
func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("lbd $ ")
		scanner.Scan()
		text := scanner.Text()
		if text == "e" || text == "exit" {
			break
		}
		l := lexer.NewLexer(lexer.StreamFromString(text))
		l.Start()
		for t := range l.NextToken {
			fmt.Println(t)
		}
		fmt.Println()
	}
}
