package parseerror

import (
	"fmt"
	"lo/token"
	"os"
)

// HadError stops code execution when we encounter a known error
var HadError = false
var HadRunTimeError = false

// SyntaxError describes a syntactic error on a given line
type SyntaxError struct {
	Token   token.Token
	Message string
}

func (e *SyntaxError) Error() string {
	HadError = true
	return MakeError(e.Token, e.Message)
}

// ParseError is thrown when an error is encoutered during token parsing
type ParseError struct {
	Token   token.Token
	Message string
}

func (e *ParseError) Error() string {
	HadError = true
	return MakeError(e.Token, e.Message)
}

// RunTimeError occured when the expression values were being evaluated
type RunTimeError struct {
	Token   token.Token
	Message string
}

func (e *RunTimeError) Error() string {
	HadRunTimeError = true
	fmt.Fprintf(os.Stderr, "[line %d] RunTime Error: %s: %s\n", e.Token.Line, e.Token.Lexeme, e.Message)
	return MakeError(e.Token, e.Message)
}

// LogError reports an error
func LogError(err error) error {
	return fmt.Errorf("%v", err)
}

// MakeError shows a parsing error as a string
func MakeError(t token.Token, message string) string {
	if t.Type == token.EOF {
		return fmt.Sprintf("[line %v] Error at end: %s", t.Line, message)
	}
	return fmt.Sprintf("[line %v] Error at '%s': %s", t.Line, t.Lexeme, message)
}

// reportError shows an error location on the stderr
func reportError(line int, where string, message string) {
	HadError = true
	fmt.Fprintf(os.Stderr, "[line %d] Error: %s: %s\n", line, where, message)
}
