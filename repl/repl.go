package repl

import (
	"fmt"
	"io"
	"os"

	"bufio"

	"github.com/BenchR267/lbd/lexer"
)

const prefix = "lbd $ "

var reader io.Reader = os.Stdin
var writer io.Writer = os.Stdout

// Start is starting the interactive REPL (currently just printing out tokens)
func Start() {

	lex := lexer.New()

	scanner := bufio.NewScanner(reader)

	for {
		fmt.Fprint(writer, prefix)
		if scanned := scanner.Scan(); !scanned {
			break
		}
		text := scanner.Text()

		input := lexer.StreamFromString(text)
		lex.Start(input)

		for t := range lex.NextToken {
			fmt.Fprintln(writer, t)
		}

		fmt.Fprintln(writer)
	}
}
