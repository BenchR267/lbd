package repl

import (
	"fmt"
	"io"
	"os"

	"bufio"

	"time"

	"github.com/BenchR267/lbd/lexer"
)

var reader io.Reader = os.Stdin
var writer io.Writer = os.Stdout

// Start is starting the interactive REPL (currently just printing out tokens)
func Start() {

	c := make(chan rune)
	l := lexer.NewLexer(c)
	l.Start()

	scanner := bufio.NewScanner(reader)
	for {
		fmt.Fprint(writer, "lbd $ ")
		scanner.Scan()
		text := scanner.Text()
		if text == "e" || text == "exit" {
			close(c)
			break
		}

		done := make(chan struct{})

		go func() {
			for _, r := range text {
				c <- r
			}
			c <- '\n'
			time.Sleep(time.Millisecond * 1)
			done <- struct{}{}
		}()

	waitLoop:
		for {
			select {
			case t := <-l.NextToken:
				fmt.Fprintln(writer, t)
			case <-done:
				break waitLoop
			}
		}

		fmt.Fprintln(writer)
	}
}
