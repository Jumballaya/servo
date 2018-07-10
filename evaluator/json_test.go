package evaluator

import (
	"github.com/jumballaya/servo/object"

	"testing"
)

func TestJSONBuiltinFunction(t *testing.T) {
	input := `
let x = '{"foo":5, "bar":7, "example": 12}';
let c = json(x);
c;
`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("Eval didn't return Hash. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.String{Value: "foo"}).HashKey():     5,
		(&object.String{Value: "bar"}).HashKey():     7,
		(&object.String{Value: "example"}).HashKey(): 12,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}

		testIntegerObject(t, pair.Value, expectedValue)
	}
}
