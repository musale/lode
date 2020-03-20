package parseerror

import (
	"fmt"
	"os"
)

// HadError stops code execution when we encounter a known error
var HadError = false

// SyntaxError describes a syntactic error on a given line
type SyntaxError struct {
	line    int
	message string
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("Error on line %d: %s", e.line, e.message)
}

// LogError reports an error
func LogError(err error) {
	HadError = true
	fmt.Fprintf(os.Stderr, "%v\n", err)
}
