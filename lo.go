package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"lo/ast"
	"lo/parser"
	"lo/scanner"
	"os"
)

// Lox language
type Lox struct {
	HadRunTimeError bool
	HadError        bool
	Interpreter     *ast.Interpreter
}

// NewLox instance
func NewLox() *Lox {
	return &Lox{HadRunTimeError: false, HadError: false, Interpreter: ast.NewInterpreter()}
}

// Read a lox filePath and load the content to the run() function
func (l *Lox) runFile(fileName string) {
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		l.HadError = true
	}
	l.run(string(fileData))
	if l.HadError {
		os.Exit(65)
	}

	if l.HadRunTimeError {
		os.Exit(70)
	}
}

// runPrompt creates a CLI that loads lox content
func (l *Lox) runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	if l.HadError {
		os.Exit(65)
	}

	if l.HadRunTimeError {
		os.Exit(70)
	}

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if line == "exit\n" || line == "quit\n" {
			fmt.Println("Exiting Lox REPL...")
			os.Exit(0)
		}
		l.run(line)

	}
}

// run interprets lox content
func (l *Lox) run(srcData string) {
	scanner := scanner.NewScanner(srcData)
	tokens := scanner.ScanTokens()
	p := parser.NewParser(tokens)
	expressions, err := p.Parse()
	if err != nil {
		l.HadError = true
		fmt.Println(err)
		return
	}
	if l.HadError {
		return
	}
	l.Interpreter.Interpret(expressions)
}

func main() {
	flag.String("file", "", "the file path to execute")
	flag.Parse()

	args := flag.Args()

	if len(args) > 1 {
		fmt.Println("Usage: ./lo [filePath]")
		os.Exit(64) // The command was used incorrectly
	} else {
		l := NewLox()
		if len(args) == 1 {
			l.runFile(args[0])
		} else {
			l.runPrompt()
		}
	}

}
