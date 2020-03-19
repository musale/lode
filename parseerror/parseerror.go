package parseerror

import "fmt"

// HadError stops code execution when we encounter a known error
var HadError = false

// SyntaxError describes a syntactic error on a given line
type SyntaxError struct {
	Line    int
	message string
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("Error on line %s: %s", e.Line, e.message)
}
