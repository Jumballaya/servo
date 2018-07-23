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

func TestNewClassInstance(t *testing.T) {
	tests := []struct {
		input string
	}{
		{`class Example {}; let x = new Example()`},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		i, ok := evaluated.(*object.Instance)
		if !ok {
			t.Fatalf("object is not Instance. got=%T (%+v)", evaluated, evaluated)
		}

		if i.Inspect() != "instance of Example" {
			t.Fatalf("instance of wrong class. Got: %s", i.Inspect())
		}
	}
}

func TestClassConstructor(t *testing.T) {
	input := `
class Example {
	let name = "";
	let constructor = fn(name) {
		this.name = name;
	}
	let greet = fn() {
		"Hello " + this.name;
	};
};
let e = new Example("Name");`

	evaluated := testEval(input)
	i, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object from e.greet() is not string. got=%T (%+v)", evaluated, evaluated)
	}

	if i.Value != "Hello Name" {
		t.Fatalf("class fields not being set in constructor")
	}
}

func TestClassConstructorNewAttribute(t *testing.T) {
	input := `
class Example {
	let constructor = fn(name) {
		this.name = name;
	}
	let greet = fn() {
		"Hello " + this.name;
	};
};
let e = new Example("Name");
e.greet();`

	evaluated := testEval(input)
	i, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object from e.greet() is not string. got=%T (%+v)", evaluated, evaluated)
	}

	if i.Value != "Hello Name" {
		t.Fatalf("class fields not being set in constructor")
	}
}

func TestClassInheritance(t *testing.T) {
	input := `
class Example {
	let constructor = fn(name) {
		this.name = name;
	}
	let greet = fn() {
		"Hello " + this.name;
	};
};

class Greeter::Example {
	let constructor = fn(name) {
		super(name);
	}

	let welcome = fn() {
		"Welcome " + this.name;
	}
};

let g = new Greeter("Name");
g.greet();
g.welcome();
`
	testEval(input)
	t.Fatalf("Not tested")
}