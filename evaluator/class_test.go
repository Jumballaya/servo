package evaluator

import (
	"fmt"
	"testing"

	"github.com/jumballaya/servo/object"
)

func TestClassLiteral(t *testing.T) {
	tests := []struct {
		input string
	}{
		{`class Example {}`},
		{`class Base {};
class Example::Base {
	let add = fn(x, y) { x + y; };
	let name = "String";
}`},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		c, ok := evaluated.(*object.Class)
		if !ok {
			t.Fatalf("object is not Class. got=%T (%+v)", evaluated, evaluated)
		}

		fmt.Println(c)
	}
}
