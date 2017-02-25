package repl

import (
	"fmt"
	"io"
	"os"

	"bufio"

	"github.com/BenchR267/lbd/lexer"
)

var reader io.Reader = os.Stdin
var writer io.Writer = os.Stdout

// Start is starting the interactive REPL (currently just printing out tokens)
func Start() {
	scanner := bufio.NewScanner(reader)
	for {
		fmt.Fprint(writer, "lbd $ ")
		scanner.Scan()
		text := scanner.Text()
		if text == "e" || text == "exit" {
			break
		}
		l := lexer.NewLexer(lexer.StreamFromString(text))
		l.Start()
		for t := range l.NextToken {
			fmt.Fprintln(writer, t)
		}
		fmt.Fprintln(writer)
	}
}
