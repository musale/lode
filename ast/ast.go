package ast

import (
	"fmt"
	"lo/token"
	"strings"
)

// Expr is the base of all expressions
type Expr interface {
	Accept(i Interpreter) interface{}
}

// AssignExpr defines = operation
type AssignExpr struct {
	Name  token.Token
	Value Expr
}

// Accept ...
func (t *AssignExpr) Accept(i Interpreter) interface{} {
	return i.VisitAssignExpression(t)
}

func (t *AssignExpr) String() string {
	var sb strings.Builder
	sb.WriteString(t.Name.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(fmt.Sprintf("%s", t.Value))
	return sb.String()
}

// BinaryExpr ...
type BinaryExpr struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

// Accept ...
func (t *BinaryExpr) Accept(i Interpreter) interface{} {
	return i.VisitBinaryExpression(t)
}

func (t *BinaryExpr) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(t.Operator.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(fmt.Sprintf("%s", t.Left))
	sb.WriteString(" ")
	sb.WriteString(fmt.Sprintf("%s", t.Right))
	sb.WriteString(")")
	return sb.String()
}

// CallExpr ...
type CallExpr struct {
	Callee    Expr
	Paren     token.Token
	Arguments []Expr
}

// Accept ...
func (c *CallExpr) Accept(i Interpreter) interface{} {
	return i.VisitCallExpression(c)
}

// String prints the call operator
func (c *CallExpr) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("call")
	sb.WriteString(" ")
	sb.WriteString(fmt.Sprint(c.Callee))
	sb.WriteString(" ")
	for _, e := range c.Arguments {
		sb.WriteString(fmt.Sprint(e))
		sb.WriteString(" ")
	}
	sb.WriteString(")")
	return sb.String()
}

// GetExpr defines a property access functionality
type GetExpr struct {
	Expression Expr
	Name       token.Token
}

// Accept ...
func (g *GetExpr) Accept(i Interpreter) interface{} {
	return i.VisitGetExpression(g)
}

// String pretty prints the class
func (g *GetExpr) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(".")
	sb.WriteString(" ")
	sb.WriteString(fmt.Sprintf("%s", g.Expression))
	sb.WriteString(" ")
	sb.WriteString(g.Name.Lexeme)
	sb.WriteString(")")
	return sb.String()
}

// GroupExpr defines a property access functionality
type GroupExpr struct {
	Expression Expr
}

// Accept ...
func (t *GroupExpr) Accept(i Interpreter) interface{} {
	return i.VisitGroupExpression(t)
}

func (t *GroupExpr) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(fmt.Sprintf("%s", t.Expression))
	sb.WriteString(")")
	return sb.String()
}

// LiteralExpr defines a property access functionality
type LiteralExpr struct {
	Object interface{}
}

// Accept ...
func (t *LiteralExpr) Accept(i Interpreter) interface{} {
	return i.VisitLiteralExpression(t)
}

func (t *LiteralExpr) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprint(t.Object))
	return sb.String()
}

// LogicalExpr defines a property access functionality
type LogicalExpr struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

// Accept ...
func (l *LogicalExpr) Accept(i Interpreter) interface{} {
	return i.VisitLogicalExpression(l)
}

// String pretty prints the unary operator
func (l *LogicalExpr) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(l.Operator.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(fmt.Sprint(l.Left))
	sb.WriteString(" ")
	sb.WriteString(fmt.Sprint(l.Right))
	sb.WriteString(")")
	return sb.String()
}

// SetExpr defines a property access functionality
type SetExpr struct {
	Object Expr
	Name   token.Token
	Value  Expr
}

// Accept ...
func (s *SetExpr) Accept(i Interpreter) interface{} {
	return i.VisitSetExpression(s)
}

// String pretty prints the setter
func (s *SetExpr) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("set")
	sb.WriteString(" ")
	sb.WriteString(fmt.Sprint(s.Object))
	sb.WriteString(" ")
	sb.WriteString(s.Name.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(fmt.Sprint(s.Value))
	sb.WriteString(")")
	return sb.String()
}

// ThisExpr defines a property access functionality
type ThisExpr struct {
	Keyword token.Token
}

// Accept ...
func (t *ThisExpr) Accept(i Interpreter) interface{} {
	return i.VisitThisExpression(t)
}

func (t *ThisExpr) String() string {
	return fmt.Sprint(t.Keyword.Lexeme)
}

// UnaryExpr defines a property access functionality
type UnaryExpr struct {
	Operator token.Token
	Right    Expr
}

// Accept ...
func (t *UnaryExpr) Accept(i Interpreter) interface{} {
	return i.VisitUnaryExpression(t)
}

func (t *UnaryExpr) String() string {
	var sb strings.Builder
	sb.WriteString(t.Operator.Lexeme)
	sb.WriteString(fmt.Sprint(t.Right))
	return sb.String()
}

// VariableExpr defines a property access functionality
type VariableExpr struct {
	Name token.Token
}

// Accept ...
func (t *VariableExpr) Accept(i Interpreter) interface{} {
	return i.VisitVariableExpression(t)
}

func (t *VariableExpr) String() string {
	return fmt.Sprint((t.Name.Lexeme))
}
