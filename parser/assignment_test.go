package parser

import (
	"testing"

	"github.com/jumballaya/servo/ast"
	"github.com/jumballaya/servo/lexer"
)

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let y = true;", "y", true},
		{"let foobar = y;", "foobar", "y"},
		{"let none = null;", "none", nil},
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

		stmt := program.Statements[0]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

		val := stmt.(*ast.LetStatement).Value
		if !testLiteralExpression(t, val, tt.expectedValue) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"return 5;", 5},
		{"return true;", true},
		{"return foobar;", "foobar"},
		{"return null;", nil},
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

		stmt := program.Statements[0]
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Fatalf("stmt not *ast.returnStatement. got=%T", stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Fatalf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
		if testLiteralExpression(t, returnStmt.ReturnValue, tt.expectedValue) {
			return
		}
	}
}

func TestParsingImports(t *testing.T) {
	tests := []string{
		"import map from 'Array';",
		"import asd from 'Array';",
		"import nums from './.demos/arrays.svo';",
		"import asd from './asdasd.svo';",
	}

	for _, input := range tests {
		l := lexer.New(input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		stmt := program.Statements[0].(*ast.ExpressionStatement)

		_, ok := stmt.Expression.(*ast.ImportExpression)
		if !ok {
			t.Fatalf("exp is not ast.ImportExpression. got=%T", stmt.Expression)
		}
	}

}

func TestReassignExpression(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
	}{
		{"let x = 5; x += 5; x;", "x"},
		{"let x = 10; x -= 5; x;", "x"},
		{"let x = 5; x *= 5; x;", "x"},
		{"let x = 10; x /= 2; x;", "x"},
		{`let x = "hello"; x += " world"; x;`, "x"},
		{`let x = 5; x = "hello"; x`, "x"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 3 {
			t.Fatalf("program.Statements does not contain 3 statements. got=%d",
				len(program.Statements))
		}

		letStmt := program.Statements[0]
		if !testLetStatement(t, letStmt, tt.expectedIdentifier) {
			return
		}

		stmt := program.Statements[2]
		val := stmt.(*ast.ExpressionStatement).Expression
		if !testIdentifier(t, val, tt.expectedIdentifier) {
			return
		}
	}
}
