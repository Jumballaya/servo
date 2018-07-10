package evaluator

import (
	"strings"

	"github.com/jumballaya/servo/lexer"
	"github.com/jumballaya/servo/object"
	"github.com/jumballaya/servo/parser"
)

func evalJSONExpression() object.Object {
	return &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. Got: %d. Want: 1", len(args))
			}

			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `json` must be STRING, got %s", args[0].Type())
			}

			requiredCode := string(args[0].Inspect())
			env := object.NewEnvironment()
			l := lexer.New(requiredCode)
			p := parser.New(l)

			program := p.ParseProgram()

			if len(p.Errors()) != 0 {
				return newError(strings.Join(p.Errors(), "\n"))
			}

			return Eval(program, env)
		},
	}
}
