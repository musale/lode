package scanner

import (
	"lo/token"
	"testing"
)

func TestScanTokens(t *testing.T) {
	source := `/* Testing comments */
	var num = 1+2+9.22
	`

	scanner := New(source)
	tokens := scanner.ScanTokens()

	testCases := []struct {
		expectedType   token.Type
		expectedLexeme string
	}{
		{token.COMMENT, "/* Testing comments */"},
		{token.VAR, "var"}, {token.IDENTIFIER, "num"},
		{token.EQUAL, "="}, {token.NUMBER, "1"}, {token.PLUS, "+"},
		{token.NUMBER, "2"}, {token.PLUS, "+"}, {token.NUMBER, "9.22"},
	}

	if len(testCases) != len(tokens)-1 {
		t.Fatalf("expected to run %d testCases but got %d", len(tokens)-1, len(testCases))
	}

	for i, tt := range testCases {
		if tt.expectedType != tokens[i].Type {
			t.Fatalf("[test %d] - wrong token Type. Expected %q, got %q", i, tokens[i].Type, tt.expectedType)
		}
		if tt.expectedLexeme != tokens[i].Lexeme {
			t.Fatalf("[test %d] - wrong token Lexeme. Expected %q, got %q", i, tokens[i].Lexeme, tt.expectedLexeme)
		}
	}
}
