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

// NewParser creates a new parser
func NewParser(tokens []token.Token) *Parser {
	return &Parser{tokens, 0, false}
}

// Parse an expression
func (p *Parser) Parse() ([]ast.Stmt, error) {
	stmts := make([]ast.Stmt, 0)
	for !p.isAtEnd() {
		stmt, err := p.declaration()
		if err != nil {
			return nil, err
		}
		stmts = append(stmts, stmt)
	}
	return stmts, nil
}

// declaration repeatedly gets called when parsing a series of
// statements in a block
func (p *Parser) declaration() (ast.Stmt, error) {
	if p.match(token.VAR) {
		decl, err := p.varDeclaration()
		if err != nil {
			return nil, err
		}
		return decl, nil
	}
	stmt, err := p.statement()
	if e, ok := err.(*parseerror.ParseError); ok {
		p.synchronize()
		return nil, e
	}
	return stmt, nil
}

// statement determines the specific statement rule matched
// by looking at the token
func (p *Parser) statement() (ast.Stmt, error) {
	if p.match(token.PRINT) {
		stmt, err := p.printStatement()
		if err != nil {
			return nil, err
		}
		return stmt, nil
	}
	expr, err := p.expressionStatement()
	if err != nil {
		return nil, err
	}
	return expr, nil
}

// varDeclaration
func (p *Parser) varDeclaration() (ast.Stmt, error) {
	typ, err := p.consume(token.IDENTIFIER, "Expected a variable name.")
	if err != nil {
		return nil, err
	}

	var initializer ast.Expr
	if p.match(token.EQUAL) {
		initializer, err = p.expression()
		if err != nil {
			return nil, err
		}
	}
	p.consume(token.SEMICOLON, "Expected a ';' after a variable declaration")
	return &ast.VarStmt{Name: typ, Initializer: initializer}, nil
}

// printStatement ...
func (p *Parser) printStatement() (ast.Stmt, error) {
	value, err := p.expression()
	if err != nil {
		return nil, err
	}
	p.consume(token.SEMICOLON, "Expected ';' after value.")
	return &ast.PrintStmt{Expression: value}, nil
}

// expressionStatement ...
func (p *Parser) expressionStatement() (ast.Stmt, error) {
	value, err := p.expression()
	if err != nil {
		return nil, err
	}
	p.consume(token.SEMICOLON, "Expected ';' after value.")
	return &ast.ExpressionStmt{Expression: value}, nil
}

// expression expands to equality rule
func (p *Parser) expression() (ast.Expr, error) {
	// expr, err := p.equality()
	// if err != nil {
	// 	return nil, err
	// }
	// return expr, nil
	return p.assignment()
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
		expr = &ast.BinaryExpr{Left: expr, Operator: operator, Right: right}
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
		expr = &ast.BinaryExpr{Left: expr, Operator: operator, Right: right}
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
		expr = &ast.BinaryExpr{Left: expr, Operator: operator, Right: right}
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
		expr = &ast.BinaryExpr{Left: expr, Operator: operator, Right: right}
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
		return &ast.UnaryExpr{Operator: operator, Right: right}, nil
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
		return &ast.LiteralExpr{Object: false}, nil
	}
	if p.match(token.TRUE) {
		return &ast.LiteralExpr{Object: true}, nil
	}
	if p.match(token.NIL) {
		return &ast.LiteralExpr{Object: nil}, nil
	}

	if p.match(token.NUMBER, token.STRING) {
		return &ast.LiteralExpr{Object: p.previous().Literal}, nil
	}
	if p.match(token.LEFTPAREN) {
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		p.consume(token.RIGHTPAREN, "Expect ')' after expression.")
		return &ast.GroupExpr{Expression: expr}, nil
	}
	if p.match(token.IDENTIFIER) {
		return &ast.VariableExpr{Name: p.previous()}, nil
	}
	return nil, &parseerror.ParseError{Token: p.peek(), Message: "Expected an expression"}
}

// consume takes in the tokens until a check to stop is reached. i.e. when
// getting the tokens inside brackets
func (p *Parser) consume(typ token.Type, message string) (token.Token, error) {
	if p.check(typ) {
		return p.advance(), nil
	}
	return p.previous(), &parseerror.ParseError{Token: p.peek(), Message: message}
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

func (p *Parser) assignment() (ast.Expr, error) {
	expr, err := p.equality()
	if err != nil {
		return nil, err
	}
	if p.match(token.EQUAL) {
		equals := p.previous()
		value, err := p.assignment()
		if err != nil {
			return nil, err
		}
		e, ok := expr.(*ast.VariableExpr)
		if ok {
			return &ast.AssignExpr{Name: e.Name, Value: value}, nil
		}
		parseerror.ReportError(equals.Line, "Invalid assignment target.")
	}
	return expr, nil
}
