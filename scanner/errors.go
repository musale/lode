package scanner

import (
	"fmt"
)

// UnexpectedCharacterError gives the error of an unexpected character in source
type UnexpectedCharacterError struct {
	line      int
	character string
}

func (e UnexpectedCharacterError) Error() string {
	return fmt.Sprintf("Unexpected character %s on line %d\n", e.character, e.line)
}

// UnterminatedStringError raised when a string is not closed with a double quote
type UnterminatedStringError struct {
	line int
}

func (e UnterminatedStringError) Error() string {
	return fmt.Sprintf("Unterminated string on col %d\n", e.line)
}

// UnterminatedCommentError raised when a string is not closed with a double quote
type UnterminatedCommentError struct {
	line int
}

func (e UnterminatedCommentError) Error() string {
	return fmt.Sprintf("Unterminated comment on col %d\n", e.line)
}
