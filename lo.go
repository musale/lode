package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"lo/parseerror"
	"lo/scanner"
	"os"
)

// Read a lox filePath and load the content to the run() function
func runFile(fileName string) {
	fileData, err := ioutil.ReadFile(fileName)
	checkError(err)

	run(string(fileData))
}

// Create a CLI that loads lox content
func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	if parseerror.HadError {
		os.Exit(65)
	}

	for {
		fmt.Print("> ")
		inputData, err := reader.ReadString('\n')
		checkError(err)
		run(inputData)
	}
}

// Interpret lox content
func run(fileData string) {
	scanner := scanner.New(fileData)
	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}

// Check if an error has occured and panic
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.String("file", "", "the file path to execute")
	flag.Parse()

	args := flag.Args()

	if len(args) > 1 {
		fmt.Println("Usage: ./lo [filePath]")
		os.Exit(64) // The command was used incorrectly
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}

}
