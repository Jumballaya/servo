package object

import (
	"github.com/jumballaya/servo/ast"
)

var AppObject = &Class{
	Name:   "App",
	Fields: buildAppFields(),
}

func buildAppFields() []*ast.LetStatement {
	var fields []*ast.LetStatement
	return fields
}
