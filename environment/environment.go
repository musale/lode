package environment

import (
	"fmt"
	"lo/parseerror"
	"lo/token"
)

// Environment has a map of variable names to their accompanying
// values
type Environment struct {
	Values map[string]interface{}
}

// NewEnvironment creates a new instance for Environment
func NewEnvironment() Environment {
	values := make(map[string]interface{})
	return Environment{Values: values}
}

// Define binds a variable to a value
func (e *Environment) Define(name string, value interface{}) {
	e.Values[name] = value
}

// Get retrieves a variable value from the environment
func (e *Environment) Get(t token.Token) interface{} {
	if e.Values[t.Lexeme] != nil {
		return e.Values[t.Lexeme]
	}
	return &parseerror.RunTimeError{Token: t, Message: fmt.Sprintf("Undefined variable '%s'.", t.Lexeme)}
}
