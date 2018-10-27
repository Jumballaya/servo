package repl

import (
	"fmt"
	"io"

	"github.com/jumballaya/servo/evaluator"
	"github.com/jumballaya/servo/lexer"
	"github.com/jumballaya/servo/object"
	"github.com/jumballaya/servo/parser"

	prompt "github.com/c-bata/go-prompt"
)

const PROMPT = ">> "

type Config struct {
	Verbose bool
}

var env *object.Environment

func Start(in io.Reader, out io.Writer, config *Config) {
	env = object.NewEnvironment()
	p := prompt.New(exec, completer, prompt.OptionPrefix(">> "))
	p.Run()
}

func Run(input string, out io.Writer, config *Config) {
	//env := stdlib.NewEnvironmentWithLib()
	env := object.NewEnvironment()
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
	}

	evaluated := evaluator.Eval(program, env)

	if config.Verbose {
		if evaluated != nil {
			fmt.Fprintf(out, evaluated.Inspect())
			fmt.Fprintf(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	fmt.Fprintf(out, "Woops! We ran into some issues!\n")
	fmt.Fprintf(out, " parser errors:\n")
	for _, msg := range errors {
		fmt.Fprintf(out, "\t"+msg+"\n")
	}
}
