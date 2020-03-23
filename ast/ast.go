package ast

import (
	"lo/token"
)

// Expr is the base of all expressions
type Expr interface {
	Accept(i Interpreter) interface{}
}

// AssignExpr defines = operation
type AssignExpr struct {
	name  token.Token
	value Expr
}

// Accept ...
func (t AssignExpr) Accept(i Interpreter) interface{} {
	return i.VisitAssignExpression(t)
}

// BinaryExpr ...
type BinaryExpr struct {
	left     Expr
	operator token.Token
	right    Expr
}

// Accept ...
func (t *BinaryExpr) Accept(i Interpreter) interface{} {
	return i.VisitBinaryExpression(t)
}

// CallExpr ...
type CallExpr struct {
	callee    Expr
	paren     token.Token
	arguments []Expr
}

// Accept ...
func (t *CallExpr) Accept(i Interpreter) interface{} {
	return i.VisitCallExpression(t)
}

// GetExpr defines a property access functionality
type GetExpr struct {
	object Expr
	name   token.Token
}

// Accept ...
func (t *GetExpr) Accept(i Interpreter) interface{} {
	return i.VisitGetExpression(t)
}

// GroupExpr defines a property access functionality
type GroupExpr struct {
	expression Expr
}

// Accept ...
func (t *GroupExpr) Accept(i Interpreter) interface{} {
	return i.VisitGroupExpression(t)
}

// LiteralExpr defines a property access functionality
type LiteralExpr struct {
	object interface{}
}

// Accept ...
func (t *LiteralExpr) Accept(i Interpreter) interface{} {
	return i.VisitLiteralExpression(t)
}

// LogicalExpr defines a property access functionality
type LogicalExpr struct {
	left     Expr
	operator token.Token
	right    Expr
}

// Accept ...
func (t *LogicalExpr) Accept(i Interpreter) interface{} {
	return i.VisitLogicalExpression(t)
}

// SetExpr defines a property access functionality
type SetExpr struct {
	object Expr
	name   token.Token
	value  Expr
}

// Accept ...
func (t *SetExpr) Accept(i Interpreter) interface{} {
	return i.VisitSetExpression(t)
}

// ThisExpr defines a property access functionality
type ThisExpr struct {
	keyword token.Token
}

// Accept ...
func (t *ThisExpr) Accept(i Interpreter) interface{} {
	return i.VisitThisExpression(t)
}

// UnaryExpr defines a property access functionality
type UnaryExpr struct {
	operator token.Token
	right    Expr
}

// Accept ...
func (t *UnaryExpr) Accept(i Interpreter) interface{} {
	return i.VisitUnaryExpression(t)
}

// VariableExpr defines a property access functionality
type VariableExpr struct {
	name token.Token
}

// Accept ...
func (t *VariableExpr) Accept(i Interpreter) interface{} {
	return i.VisitVariableExpression(t)
}
