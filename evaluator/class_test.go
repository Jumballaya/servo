package evaluator

import (
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
		_, ok := evaluated.(*object.Class)
		if !ok {
			t.Fatalf("object is not Class. got=%T (%+v)", evaluated, evaluated)
		}
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
	s, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object from e.greet() is not string. got=%T (%+v)", evaluated, evaluated)
	}

	if s.Value != "Hello Name" {
		t.Fatalf("class fields not being set in constructor. Expected: %s, Got: %s", "Hello Name", s.Value)
	}
}

func TestClassInheritance(t *testing.T) {
	input := `
class Example {
	let constructor = fn(firstname, lastname) {
		this.firstname = firstname;
		this.lastname = lastname;
	}
	let greet = fn() {
		"Hello " + this.firstname + " " + this.lastname;
	};
};

class Greeter::Example {
	let constructor = fn(firstname, lastname) {
		super(firstname, lastname);
	}

	let welcome = fn() {
		"Welcome " + this.firstname + " " + this.lastname;
	}
};

let g = new Greeter("Firstname", "Lastname");
g.welcome();
`
	evaluated := testEval(input)
	s, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object from e.greet() is not string. got=%T (%+v)", evaluated, evaluated)
	}

	if s.Value != "Welcome Firstname Lastname" {
		t.Fatalf("class fields not being set in constructor. Expected: %s, Got: %s", "Hello Name", s.Value)
	}
}
