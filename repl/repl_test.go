package repl

import (
	"bytes"
	"strings"
	"testing"
)

func TestStart(t *testing.T) {
	reader = strings.NewReader("a=5\nb=3")
	buf := new(bytes.Buffer)
	writer = buf

	Start()

	s := buf.String()

	if !strings.HasPrefix(s, "lbd $ ") {
		t.Error("Expected output to have prefix 'lbd $ '")
	}

	lines := strings.Split(s, "\n")
	if len(lines) != 9 {
		t.Errorf("Expected to get 5 output lines, but got %d instead.", len(lines))
	}
}
