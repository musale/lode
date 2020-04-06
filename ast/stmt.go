package ast

import "go/token"

// Stmt interface for statements
type Stmt interface {
	Accept(i Interpreter) interface{}
}

// PrintStmt ...
type PrintStmt struct {
	Expression Expr
}

// Accept visits the PrintStmt
func (stmt *PrintStmt) Accept(i Interpreter) interface{} {

	return i.VisitPrintStmt(stmt)
}

// ExpressionStmt ...
type ExpressionStmt struct {
	Expression Expr
}

// Accept visits the ExpressionStmt
func (stmt *ExpressionStmt) Accept(i Interpreter) interface{} {
	return i.VisitExpressionStmt(stmt)
}

// VarStmt statement
type VarStmt struct {
	Name        token.Token
	Initializer Expr
}

// Accept visits the VarStmt
func (stmt *VarStmt) Accept(i Interpreter) interface{} {
	return i.VisitVarStmt(stmt)

}
