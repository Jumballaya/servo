package stdlib

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/jumballaya/servo/evaluator"
	"github.com/jumballaya/servo/lexer"
	"github.com/jumballaya/servo/object"
	"github.com/jumballaya/servo/parser"
)

var Libs = []string{
	"Array",
	"String",
}

func NewEnvironmentWithLib() *object.Environment {
	newEnv := object.NewEnvironment()
	for _, lib := range Libs {
		// Figure out a common place to put these files
		currentFile := "./stdlib/.svo/" + lib + ".svo"
		dir, err := filepath.Abs(currentFile)
		if err != nil {
			log.Fatal(err)
		}
		file, err := ioutil.ReadFile(dir)
		if err != nil {
			fmt.Println(err.Error())
			log.Fatal(err)
		}

		requiredCode := string(file[:])
		env := object.NewEnvironment()
		l := lexer.New(requiredCode)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			log.Fatal(p.Errors())
		}

		embedded := evaluator.Eval(program, env)
		newEnv.Set(lib, embedded)

	}

	return newEnv
}
