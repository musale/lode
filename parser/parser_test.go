package parser

import (
	"fmt"
	"lo/scanner"
	"testing"
)

func TestParser(t *testing.T) {
	source := `1+2+9.22;`
	sc := scanner.New(source)
	tokens := sc.ScanTokens()
	pa := New(tokens)
	expression, err := pa.parse()
	if err != nil {
		t.Errorf("%s", err)
	}
	expected := "(+ (+ 1 2) 9.22)"

	if fmt.Sprintf("%s", expression) != expected {
		t.Errorf("expected %s but got %s", expected, expression)
	}
}
