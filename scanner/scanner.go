package scanner

import (
	"lo/parseerror"
	"lo/token"
	"strconv"
)

var keyWords = map[string]token.Type{
	"and":    token.AND,
	"class":  token.CLASS,
	"else":   token.ELSE,
	"false":  token.FALSE,
	"for":    token.FOR,
	"fun":    token.FUN,
	"if":     token.IF,
	"nil":    token.NIL,
	"or":     token.OR,
	"print":  token.PRINT,
	"return": token.RETURN,
	"super":  token.SUPER,
	"this":   token.THIS,
	"true":   token.TRUE,
	"var":    token.VAR,
	"while":  token.WHILE,
}

// Scanner ...
type Scanner struct {
	start, current, line int
	source               string
	tokens               []token.Token
}

// NewScanner creates a new Scanner
func NewScanner(source string) *Scanner {
	return &Scanner{line: 1, source: source, tokens: make([]token.Token, 0)}
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
		if s.peekBack() != "/" || s.peek() != "/" {
			s.addToken(token.STAR)
		} else {
			for s.peek() != "\n" && !s.isAtEnd() {
				s.advance()
			}
		}
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
		if s.peek() == "*" {
			s.advance()
			s.parseComment()
		} else if s.match("/") {
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
		if s.isDigit(sourceChar) {
			s.number()
		} else if s.isAlpha(sourceChar) {
			s.identifier()
		} else {
			parseerror.LogError(UnexpectedCharacterError{line: s.line, character: sourceChar})
		}
	}
}

// parseComment sets the comment tokens
func (s *Scanner) parseComment() {
	for s.peek() != "*" && s.peekNext() != "/" && !s.isAtEnd() {
		s.line++
		s.advance()
	}
	if s.isAtEnd() {
		parseerror.LogError(UnterminatedCommentError{line: s.line})
		return
	}
	s.advance()
	s.advance()
	// string the double quotes on both ends of the string
	commentString := s.source[s.start+2 : s.current-2]
	s.addTokenWithLiteral(token.COMMENT, string(commentString))
}

// identifier sets an identifier
func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}
	var tokenType token.Type
	tokenText := string(s.source[s.start:s.current])
	if keyWords[tokenText] == "" {
		tokenType = token.IDENTIFIER
	} else {
		tokenType = keyWords[tokenText]
	}
	s.addToken(tokenType)
}

// isAlpha checks whether a character is an alphabet or _
func (s *Scanner) isAlpha(char string) bool {
	return char >= "a" && char <= "z" || char >= "A" && char <= "Z" || char == "_"
}

// isAlphaNumeric checks whether a character is an alphabet, _ or number
func (s *Scanner) isAlphaNumeric(char string) bool {
	return s.isAlpha(char) || s.isDigit(char)
}

// isDigit determines whether the character is a number between 0-9
func (s *Scanner) isDigit(char string) bool {
	return char >= "0" && char <= "9"
}

// number consumes a number literal
func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	// Check if the number is followed by a decimal and a number
	if s.peek() == "." && s.isDigit(s.peekNext()) {
		s.advance()
		for s.isDigit(s.peek()) {
			s.advance()
		}
	}
	// Maybe handle this error?
	numberValue, _ := strconv.ParseFloat(s.source[s.start:s.current], 64)
	s.addTokenWithLiteral(token.NUMBER, numberValue)
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

// peek looks ahead one character without consuming any character
func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return "EOF"
	}
	return s.currentCharacter()
}

// peekBack looks back one character without consuming any character
func (s *Scanner) peekBack() string {
	if s.isAtEnd() {
		return "SOF"
	}
	return string(s.source[s.current-1])
}

// peekNext looks ahead at the character after peek()
func (s *Scanner) peekNext() string {
	if s.current+1 >= len(s.source) {
		return "EOF"
	}
	return string(s.source[s.current+1])
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
