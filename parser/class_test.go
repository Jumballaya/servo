package parser

import (
	"testing"

	"github.com/jumballaya/servo/ast"
	"github.com/jumballaya/servo/lexer"
)

func TestClassLiterals(t *testing.T) {
	tests := []struct {
		input              string
		expectedName       string
		expectedParent     string
		expectedStatements int
	}{
		{"class Example::Parent {};", "Example", "Parent", 0},
		{`class Example {
	let name = "test";
	let add = fn(x, y) { x + y };
};`, "Example", "", 2},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		stmt := program.Statements[0].(*ast.LetStatement)
		literal, ok := stmt.Value.(*ast.ClassLiteral)
		if !ok {
			t.Fatalf("exp not *ast.ClassLiteral. Got: %T", stmt.Value)
		}

		if tt.expectedParent != literal.Parent {
			t.Fatalf("class parent is incorrect. Wanted: %s Got: %s", tt.expectedParent, literal.Parent)
		}

		if tt.expectedName != literal.Name {
			t.Fatalf("class name is incorrect. Wanted: %s Got: %s", tt.expectedName, literal.Name)
		}

		if tt.expectedStatements != (len(literal.Methods) + len(literal.Fields)) {
			t.Fatalf("class parent is incorrect. Wanted: %s Got: %s", tt.expectedParent, literal.Parent)
		}
	}
}
