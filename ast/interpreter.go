package ast

// Interpreter ...
type Interpreter struct{}

// VisitAssignExpression ...
func (i Interpreter) VisitAssignExpression(e *AssignExpr) interface{} {
	return ""
}

// VisitBinaryExpression ...
func (i Interpreter) VisitBinaryExpression(e *BinaryExpr) interface{} {
	return ""
}

// VisitCallExpression ...
func (i Interpreter) VisitCallExpression(e *CallExpr) interface{} {
	return ""
}

// VisitGetExpression ...
func (i Interpreter) VisitGetExpression(e *GetExpr) interface{} {
	return ""
}

// VisitGroupExpression ...
func (i Interpreter) VisitGroupExpression(e *GroupExpr) interface{} {
	return ""
}

// VisitLiteralExpression ...
func (i Interpreter) VisitLiteralExpression(e *LiteralExpr) interface{} {
	return ""
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
	return ""
}

// VisitVariableExpression ...
func (i Interpreter) VisitVariableExpression(e *VariableExpr) interface{} {
	return ""
}
