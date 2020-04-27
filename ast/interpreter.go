package ast

import (
	"fmt"
	"lo/environment"
	"lo/parseerror"
	"lo/token"
	"reflect"
)

// Interpreter ...
type Interpreter struct {
	Environment environment.Environment
}

// NewInterpreter creates a new interpreter
func NewInterpreter() *Interpreter {
	env := environment.NewEnvironment()
	return &Interpreter{Environment: env}
}

// Interpret the given expressions
func (i Interpreter) Interpret(stmts []Stmt) {
	// value := i.evaluate(e)
	// fmt.Println(value)
	for _, stmt := range stmts {
		i.evaluate(stmt)
	}
}

func (i *Interpreter) String() string {
	if i == nil {
		return "nil"
	}
	return i.String()
}

// VisitAssignExpression ...
func (i Interpreter) VisitAssignExpression(e *AssignExpr) interface{} {
	value := i.evaluate(e.Value)
	i.Environment.Assign(e.Name, value)
	return value
}

// VisitBinaryExpression ...
func (i Interpreter) VisitBinaryExpression(e *BinaryExpr) interface{} {
	left := i.evaluate(e.Left)
	right := i.evaluate(e.Right)

	switch e.Operator.Type {
	case token.MINUS:
		i.checkOneNumberOperand(e.Operator, right)
		return left.(float64) - right.(float64)
	case token.SLASH:
		i.checkTwoNumberOperands(e.Operator, left, right)
		return left.(float64) / right.(float64)
	case token.STAR:
		i.checkTwoNumberOperands(e.Operator, left, right)
		return left.(float64) * right.(float64)
	case token.GREATER:
		i.checkTwoNumberOperands(e.Operator, left, right)
		return left.(float64) > right.(float64)
	case token.GREATEREQUAL:
		i.checkTwoNumberOperands(e.Operator, left, right)
		return left.(float64) >= right.(float64)
	case token.LESS:
		i.checkTwoNumberOperands(e.Operator, left, right)
		return left.(float64) < right.(float64)
	case token.LESSEQUAL:
		i.checkTwoNumberOperands(e.Operator, left, right)
		return left.(float64) <= right.(float64)
	case token.BANGEQUAL:
		return !i.isEqual(left, right)
	case token.EQUALEQUAL:
		return i.isEqual(left, right)
	case token.PLUS:
		typeLeft := reflect.TypeOf(left).String()
		typeRight := reflect.TypeOf(right).String()
		if (typeLeft == "float64" || typeLeft == "float32" || typeLeft == "int") && (typeRight == "float64" || typeRight == "float32" || typeRight == "int") {
			return left.(float64) + right.(float64)
		} else if typeLeft == "string" && typeRight == "string" {
			return left.(string) + right.(string)
		}
		return &parseerror.RunTimeError{
			Token:   e.Operator,
			Message: fmt.Sprintf("Operand %v and %v must be numbers or strings", left, right)}
	}
	return nil
}

// checkOneNumberOperand if it's a float or int
func (i Interpreter) checkOneNumberOperand(operator token.Token, operand interface{}) error {
	typeOfOperand := reflect.TypeOf(operand).String()
	if typeOfOperand == "float64" || typeOfOperand == "float32" || typeOfOperand == "int" {
		return nil
	}
	return &parseerror.RunTimeError{Token: operator, Message: fmt.Sprintf("Operand %v must be a number", operand)}
}

// checkTwoNumberOperands if they are floats or ints
func (i Interpreter) checkTwoNumberOperands(operator token.Token, left interface{}, right interface{}) error {
	typeOfLeftOperand := reflect.TypeOf(left).String()
	typeOfRightOperand := reflect.TypeOf(right).String()
	if (typeOfLeftOperand == "float64" || typeOfLeftOperand == "float32" || typeOfLeftOperand == "int") &&
		(typeOfRightOperand == "float64" || typeOfRightOperand == "float32" || typeOfRightOperand == "int") {
		return nil
	}
	return &parseerror.RunTimeError{Token: operator, Message: fmt.Sprintf("Operand %v and %v must be a number", left, right)}
}

// isEqual returns true if 2 objects are the same
func (i Interpreter) isEqual(left interface{}, right interface{}) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil {
		return false
	}
	return reflect.DeepEqual(left, right)
}

// VisitCallExpression ...
func (i Interpreter) VisitCallExpression(e *CallExpr) interface{} {
	return ""
}

// VisitGetExpression ...
func (i Interpreter) VisitGetExpression(e *GetExpr) interface{} {
	return ""
}

// VisitGroupExpression resturns the result of values in parenthesis
// expression
func (i Interpreter) VisitGroupExpression(e *GroupExpr) interface{} {
	return i.evaluate(e.Expression)
}

// evaluate is a helper that revisits the interpretor
func (i Interpreter) evaluate(e Expr) interface{} {
	return e.Accept(i)
}

// VisitLiteralExpression returns the runtime value the parser took
func (i Interpreter) VisitLiteralExpression(e *LiteralExpr) interface{} {
	return e.Object
}

// VisitLogicalExpression ...
func (i Interpreter) VisitLogicalExpression(e *LogicalExpr) interface{} {
	return ""
}

// VisitSetExpression ...
func (i Interpreter) VisitSetExpression(e *SetExpr) interface{} {
	return ""
}

// VisitThisExpression ...
func (i Interpreter) VisitThisExpression(e *ThisExpr) interface{} {
	return ""
}

// VisitUnaryExpression ...
func (i Interpreter) VisitUnaryExpression(e *UnaryExpr) interface{} {
	right := i.evaluate(e.Right)
	switch e.Operator.Type {
	case token.MINUS:
		value := -right.(float64)
		return value
	case token.BANG:
		return !i.isTruthy(true)
	}
	return nil
}

// isTruthy returns false and nil object as falsey and everything else as
// truthy
func (i Interpreter) isTruthy(obj interface{}) bool {
	if obj == nil {
		return false
	}
	kind := reflect.TypeOf(obj).Kind()
	switch kind {
	case reflect.Bool:
		return obj.(bool)
	default:
		return true
	}
}

// VisitVariableExpression ...
func (i Interpreter) VisitVariableExpression(e *VariableExpr) interface{} {
	return i.Environment.Get(e.Name)
}

// VisitPrintStmt ...
func (i Interpreter) VisitPrintStmt(e *PrintStmt) interface{} {
	value := i.evaluate(e.Expression)
	fmt.Println(fmt.Sprint(value))
	return nil
}

// VisitExpressionStmt ...
func (i Interpreter) VisitExpressionStmt(e *ExpressionStmt) interface{} {
	i.evaluate(e.Expression)
	return nil
}

// VisitVarStmt ...
func (i Interpreter) VisitVarStmt(e *VarStmt) interface{} {
	var value interface{}
	if e.Initializer != nil {
		value = i.evaluate(e.Initializer)
	}
	i.Environment.Define(e.Name.Lexeme, value)
	return nil
}
