package parser

import (
	"testing"

	"github.com/BenchR267/lbd/lexer/token"
)

func TestParseError(t *testing.T) {
	t1 := token.Token{
		Raw: "foo",
	}
	t2 := token.Token{
		Raw: "bar",
	}

	e := ParseError{
		Expected: t1,
		Got:      t2,
	}

	if e.Error() != "Expected foo, but got bar." {
		t.Errorf("Error output was different than expected. (%s)", e.Error())
	}
}
