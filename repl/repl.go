package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jumballaya/servo/evaluator"
	"github.com/jumballaya/servo/lexer"
	"github.com/jumballaya/servo/parser"
	"github.com/jumballaya/servo/stdlib"
)

const PROMPT = ">> "

type History struct {
	Current int
	History []string
}

func NewHistory() *History {
	return &History{Current: 0, History: []string{""}}
}

func (h *History) Back() {
	length := len(h.History)
	if h.Current < 0 || h.Current > length-1 {
		h.Current = 0
		return
	}

	h.Current--
}

func (h *History) Forward() {
	length := len(h.History)
	if h.Current < 0 {
		h.Current = 0
		return
	}
	if h.Current > length-1 {
		h.Current = length - 1
		return
	}

	h.Current += 1
}

func (h *History) CurrentInput() string {
	return h.History[h.Current]
}

func (h *History) Insert(command string) {
	h.History[h.Current] = command
	h.Current = len(h.History) + 1
	h.History = append(h.History, "")
}

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := stdlib.NewEnvironmentWithLib()
	//history := NewHistory()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		//history.Insert(line)

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func Run(input string, out io.Writer, verbose bool) {
	env := stdlib.NewEnvironmentWithLib()
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
	}

	evaluated := evaluator.Eval(program, env)

	if verbose {
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some issues!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
