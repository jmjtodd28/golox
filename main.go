package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jmjtodd28/golox/parser"
	"github.com/jmjtodd28/golox/scanner"
)

func main() {

	args := os.Args[1:]

	if len(args) > 1 {
		log.Fatal("Usage: golox [script]")
	} else if len(args) == 1 {
		if err := runFile(args[0]); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := runPrompt(); err != nil {
			log.Fatal(err)
		}
	}
}

func runFile(script string) error {
	fileContents, err := os.ReadFile(script)
	if err != nil {
		return fmt.Errorf("Opening file: %w:", err)
	}
	if err := run(string(fileContents)); err != nil {
		return err
	}
	return nil
}

func runPrompt() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to lox REPL. ctrl-D to exit")

	for {
		fmt.Print(">>> ")
		input, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				fmt.Println("EXITING")
				break
			}
			panic("Unexpected error reading from user")
		}
		if err := run(input); err != nil {
			return err
		}
	}
	return nil
}

func run(script string) error {
	scanner := scanner.NewScanner(script)
	scanner.ScanTokens()

	for _, token := range scanner.Tokens {
		fmt.Println(token.String())
	}

	parser := parser.NewParser(scanner.Tokens)
	parser.Parse()
	return nil
}
