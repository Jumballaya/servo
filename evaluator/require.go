package evaluator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/jumballaya/servo/lexer"
	"github.com/jumballaya/servo/object"
	"github.com/jumballaya/servo/parser"
)

func evalRequireExpression() object.Object {
	return &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. Got: %d. Want: 1", len(args))
			}

			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `push` must be STRING, got %s", args[0].Type())
			}

			requiredFile := args[0].Inspect()
			currentFile := os.Args[1]
			currentDir := "./" + strings.Join(strings.Split(currentFile, "/")[:1], "/")

			dir, err := filepath.Abs(currentDir + "/" + requiredFile)
			if err != nil {
				fmt.Println(err.Error())
				return newError(err.Error())
			}

			file, err := ioutil.ReadFile(dir)
			if err != nil {
				fmt.Println(err.Error())
				return newError(err.Error())
			}

			requiredCode := string(file[:])
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
