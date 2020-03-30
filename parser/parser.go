package parser

import (
	"lo/ast"
	"lo/parseerror"
	"lo/token"
)

// Parser consumes Tokens and uses current to point to the next Token position
type Parser struct {
	tokens  []token.Token
	current int64
	inloop  bool
}

// New creates a new parser
func New(tokens []token.Token) Parser {
	return Parser{tokens, 0, false}
}

func (p *Parser) parse() (ast.Expr, error) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}
	return expr, nil
}

// expression expands to equality rule
func (p *Parser) expression() (ast.Expr, error) {
	expr, err := p.equality()
	if err != nil {
		return nil, err
	}
	return expr, nil
}

// equalilty handles the != and == expressions
func (p *Parser) equality() (ast.Expr, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}

	for p.match(token.BANGEQUAL, token.EQUALEQUAL) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{expr, operator, right}
	}
	return expr, nil
}

// match returns true if current token is any of the given types
func (p *Parser) match(types ...token.Type) bool {
	for _, typ := range types {
		if p.check(typ) {
			p.advance()
			return true
		}
	}
	return false
}

// check if the current token type is of the given type
func (p *Parser) check(typ token.Type) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == typ
}

// advance looks at the current token and moves the pointer ahead if it's
// of the type required
func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

// isAtEnd checks if you've run out of tokens to parse
func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

// peek returns the Token at the current position
func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}

// previous returns the Token recently consumed
func (p *Parser) previous() token.Token {
	return p.tokens[p.current-1]
}

// comparison handles the >, >=, < and <= expressions
func (p *Parser) comparison() (ast.Expr, error) {
	expr, err := p.addition()
	if err != nil {
		return nil, err
	}
	for p.match(token.GREATER, token.GREATEREQUAL, token.LESS, token.LESSEQUAL) {
		operator := p.previous()
		right, err := p.addition()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{expr, operator, right}
	}
	return expr, nil
}

// addition handles the - and + expressions
func (p *Parser) addition() (ast.Expr, error) {
	expr, err := p.multiplication()
	if err != nil {
		return nil, err
	}
	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right, err := p.multiplication()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{expr, operator, right}
	}
	return expr, nil
}

// multiplication handles the / and * expressions
func (p *Parser) multiplication() (ast.Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}
	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{expr, operator, right}
	}
	return expr, nil
}

// unary handles the ! and - expressions
func (p *Parser) unary() (ast.Expr, error) {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return &ast.UnaryExpr{operator, right}, nil
	}
	expr, err := p.primary()
	if err != nil {
		return nil, err
	}
	return expr, nil
}

// primary is the highest level of precedence handling the basic expressions
func (p *Parser) primary() (ast.Expr, error) {
	if p.match(token.FALSE) {
		return &ast.LiteralExpr{false}, nil
	}
	if p.match(token.TRUE) {
		return &ast.LiteralExpr{true}, nil
	}
	if p.match(token.NIL) {
		return &ast.LiteralExpr{nil}, nil
	}

	if p.match(token.NUMBER, token.STRING) {
		return &ast.LiteralExpr{p.previous().Literal}, nil
	}
	if p.match(token.LEFTPAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		p.consume(token.RIGHTPAREN, "Expect ')' after expression.")
		return &ast.GroupExpr{expr}, nil
	}
	return nil, &parseerror.ParseError{p.peek(), "Expected an expression"}
}

// consume takes in the tokens until a check to stop is reached. i.e. when
// getting the tokens inside brackets
func (p *Parser) consume(typ token.Type, message string) (token.Token, error) {
	if p.check(typ) {
		return p.advance(), nil
	}
	return p.previous(), &parseerror.ParseError{p.peek(), message}
}

// synchronize discards token until it finds a statement
func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().Type == token.SEMICOLON {
			return
		}
		switch p.peek().Type {
		case token.CLASS:
		case token.FUN:
		case token.VAR:
		case token.FOR:
		case token.IF:
		case token.WHILE:
		case token.PRINT:
		case token.RETURN:
			return
		}
		p.advance()
	}
}
