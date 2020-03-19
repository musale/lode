package scanner

import (
	"fmt"
	"lo/token"
)

// Scanner ...
type Scanner struct{}

// New creates a new Scanner
func New(source string) Scanner {
	fmt.Println(source)
	return Scanner{}
}

// ScanTokens ...
func (s Scanner) ScanTokens() []token.Token {
	return []token.Token{}
}
