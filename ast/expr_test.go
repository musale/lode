package ast

import (
	"fmt"
	"lo/token"
	"testing"
)

func TestExpr(t *testing.T) {
	left := &UnaryExpr{
		token.Token{Type: token.MINUS, Lexeme: "-", Literal: nil, Line: 1}, &LiteralExpr{123},
	}
	right := &GroupExpr{&LiteralExpr{245}}
	operator := token.Token{Type: token.PLUS, Lexeme: "+", Literal: nil, Line: 1}
	expression := &BinaryExpr{left, operator, right}
	expected := "(+ -123 (245))"

	if fmt.Sprintf("%s", expression) != expected {
		t.Errorf("expected %s but got %s", expected, expression)
	}
}
