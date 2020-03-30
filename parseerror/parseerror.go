package parseerror

import (
	"fmt"
	"lo/token"
	"os"
)

// HadError stops code execution when we encounter a known error
var HadError = false

// SyntaxError describes a syntactic error on a given line
type SyntaxError struct {
	token   token.Token
	message string
}

func (e *SyntaxError) Error() string {
	reportError(e.token.Line, e.token.Lexeme, e.message)
	return MakeError(e.token, e.message)
}

// ParseError is thrown when an error is encoutered during token parsing
type ParseError struct {
	token   token.Token
	message string
}

func (e *ParseError) Error() string {
	reportError(e.token.Line, e.token.Lexeme, e.message)
	return MakeError(e.token, e.message)
}

// LogError reports an error
func LogError(err error) {
	HadError = true
	fmt.Fprintf(os.Stderr, "%v\n", err)
}

// MakeError shows a parsing error as a string
func MakeError(token token.Token, message string) error {
	if token.Type == token.EOF {
		return fmt.Errorf("[line %v] Error at end: %s", token.Line, message)
	}
	return fmt.Errorf("[line %v] Error at '%s': %s", token.Line, token.Lexeme, message)
}

// reportError shows an error location on the stderr
func reportError(line int, where string, message string) {
	HadError = true
	fmt.Fprintf(os.Stderr, "[line %d] Error: %s: %s\n", line, where, message)
}
