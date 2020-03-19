package scanner

import (
	"lo/parseerror"
	"lo/token"
)

// Scanner ...
type Scanner struct {
	start, current, line int
	source               string
	tokens               []token.Token
}

// New creates a new Scanner
func New(source string) Scanner {
	return Scanner{line: 1, source: source, tokens: make([]token.Token, 0)}
}

// ScanTokens consumes the tokens in a source and returns them set to their types
func (s *Scanner) ScanTokens() []token.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, token.Token{Type: token.EOF})
	return s.tokens
}

// scanToken determines the type of Token and adds it to the Scanner
func (s *Scanner) scanToken() {
	sourceChar := s.advance()

	switch sourceChar {
	case "(":
		s.addToken(token.LEFTPAREN)
	case ")":
		s.addToken(token.RIGHTPAREN)
	case "{":
		s.addToken(token.LEFTBRACE)
	case "}":
		s.addToken(token.RIGHTBRACE)
	case ",":
		s.addToken(token.COMMA)
	case ".":
		s.addToken(token.DOT)
	case "-":
		s.addToken(token.MINUS)
	case "+":
		s.addToken(token.PLUS)
	case ";":
		s.addToken(token.SEMICOLON)
	case "*":
		s.addToken(token.STAR)
	case "!": // !=
		if !s.match("=") {
			s.addToken(token.BANG)
		} else {
			s.addToken(token.BANGEQUAL)
		}
	case "=": // ==
		if !s.match("=") {
			s.addToken(token.EQUAL)
		} else {
			s.addToken(token.EQUALEQUAL)
		}
	case "<": // <=
		if !s.match("=") {
			s.addToken(token.LESS)
		} else {
			s.addToken(token.LESSEQUAL)
		}
	case ">": // >=
		if !s.match("=") {
			s.addToken(token.GREATER)
		} else {
			s.addToken(token.GREATEREQUAL)
		}
	case "/":
		if s.match("/") {
			for s.peek() != "\n" && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH)
		}
	case " ":
	case "\r":
	case "\t":
		break
	case "\n":
		s.line++
	case "\"":
		s.parseString()
	default:
		parseerror.LogError(UnexpectedCharacterError{line: s.line, character: sourceChar})
	}
}

// parseString consumes a string from the opening to the closing double quote
func (s *Scanner) parseString() {
	for s.peek() != "\"" && !s.isAtEnd() {
		s.line++
		s.advance()
	}
	if s.isAtEnd() {
		parseerror.LogError(UnterminatedStringError{line: s.line})
		return
	}
	s.advance()
	// string the double quotes on both ends of the string
	stringValue := s.source[s.start+1 : s.current-1]
	s.addTokenWithLiteral(token.STRING, string(stringValue))
}

// peek looksahead without consuming any character
func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return "EOF"
	}
	return s.currentCharacter()
}

// match checks the next character in the source and determines if it is a
// one or two character token
func (s *Scanner) match(character string) bool {
	if s.isAtEnd() {
		return false
	}
	if s.currentCharacter() != character {
		return false
	}
	s.current++
	return true
}

// currentCharacter gets and returns the current character in the source
func (s *Scanner) currentCharacter() string {
	return string(s.source[s.current])
}

// advance increases the value of the current position in the source. It then
// determines the character at the current position and returns it
func (s *Scanner) advance() string {
	s.current++
	return string(s.source[s.current-1])
}

// addToken appends a new token to the scanner
func (s *Scanner) addToken(token token.Type) {
	s.addTokenWithLiteral(token, nil)
}

// addTokenWithLiteral sets a token with a Literal value i.e. string tokens
func (s *Scanner) addTokenWithLiteral(tokenType token.Type, literal interface{}) {
	lexeme := string(s.source[s.start:s.current])
	s.tokens = append(s.tokens, token.Token{Type: tokenType, Lexeme: lexeme, Line: s.line, Literal: literal})
}

// isAtEnd signals consumption of all the characters in a source
func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
