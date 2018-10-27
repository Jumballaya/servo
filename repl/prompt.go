package repl

import (
	"fmt"
	"os"

	prompt "github.com/c-bata/go-prompt"
	"github.com/jumballaya/servo/evaluator"
	"github.com/jumballaya/servo/lexer"
	"github.com/jumballaya/servo/parser"
)

func exec(line string) {
	if line == "quit" || line == "exit" || line == "q" {
		os.Exit(0)
	}

	l := lexer.New(line)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		for _, e := range p.Errors() {
			fmt.Println(e)
		}
	}
	evaluated := evaluator.Eval(program, env)
	fmt.Println(evaluated.Inspect())
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "exit", Description: "exit the repl"},
		{Text: "let", Description: "let [descriptor] = [value]"},
		{Text: "class", Description: "class [descriptor] {}"},
		{Text: "fn", Description: "fn [descriptor]([...arguments] { [...code] })"},
	}

	for k, v := range env.FullList() {
		s = append(s, prompt.Suggest{
			Text:        k,
			Description: v,
		})
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
