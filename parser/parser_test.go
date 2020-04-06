package parser

import (
	"fmt"
	"lo/scanner"
	"testing"
)

func TestParser(t *testing.T) {
	source := `1+2+9.22;`
	sc := scanner.NewScanner(source)
	tokens := sc.ScanTokens()
	pa := NewParser(tokens)
	expression, err := pa.Parse()
	if err != nil {
		t.Errorf("%s", err)
	}
	expected := "(+ (+ 1 2) 9.22)"

	if fmt.Sprintf("%s", expression) != expected {
		t.Errorf("expected %s but got %s", expected, expression)
	}
}
