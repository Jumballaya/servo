package evaluator

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/jumballaya/servo/lexer"
	"github.com/jumballaya/servo/object"
	"github.com/jumballaya/servo/parser"
)

func LoadAndEvalFile(file string) object.Object {
	requiredCode, err := LoadFile(file)
	if err != nil {
		return newError(err.Error())
	}
	env := object.NewEnvironment()
	l := lexer.New(requiredCode)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		return newError(strings.Join(p.Errors(), "\n"))
	}

	return Eval(program, env)
}

func GetObjectFromFile(file, objName string) object.Object {
	requiredCode, err := LoadFile(file)
	if err != nil {
		return newError(err.Error())
	}

	env := object.NewEnvironment()
	l := lexer.New(requiredCode)
	p := parser.New(l)
	program := p.ParseProgram()
	Eval(program, env)

	if len(p.Errors()) != 0 {
		return newError(strings.Join(p.Errors(), "\n"))
	}

	if val, ok := env.Get(objName); ok {
		return val
	}

	return newError("identifier not found %s", objName)
}

func LoadFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(data[:]), nil
}
