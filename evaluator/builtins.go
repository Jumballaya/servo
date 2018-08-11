package evaluator

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/jumballaya/servo/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. Got: %d. Want: 1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. Got: %d. Want: 1", len(args))
			}

			switch args[0].(type) {
			case *object.Array:
				arr := args[0].(*object.Array)
				if len(arr.Elements) > 0 {
					return arr.Elements[0]
				}
				return NULL
			case *object.String:
				str := args[0].(*object.String)
				if len(str.Value) > 0 {
					return &object.String{Value: strings.Split(str.Value, "")[0]}
				}
				return NULL
			default:
				return newError("argument to `first` must be ARRAY or STRING, got %s", args[0].Type())
			}
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. Got: %d. Want: 1", len(args))
			}

			switch args[0].(type) {
			case *object.Array:
				arr := args[0].(*object.Array)
				if len(arr.Elements) > 0 {
					return arr.Elements[len(arr.Elements)-1]
				}
				return NULL
			case *object.String:
				str := args[0].(*object.String)
				if len(str.Value) > 0 {
					val := strings.Split(str.Value, "")
					return &object.String{Value: val[len(val)-1]}
				}
				return NULL
			default:
				return newError("argument to `last` must be ARRAY or STRING, got %s", args[0].Type())
			}
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. Got: %d. Want: 1", len(args))
			}

			switch args[0].(type) {
			case *object.Array:
				arr := args[0].(*object.Array)
				length := len(arr.Elements)
				if length > 0 {
					newElements := make([]object.Object, length-1, length-1)
					copy(newElements, arr.Elements[1:length])
					return &object.Array{Elements: newElements}
				}
				return NULL
			case *object.String:
				str := args[0].(*object.String)
				if len(str.Value) > 0 {
					val := strings.Split(str.Value, "")
					return &object.String{Value: strings.Join(val[1:], "")}
				}
				return NULL
			default:
				return newError("argument to `rest` must be ARRAY or STRING, got %s", args[0].Type())
			}
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. Got: %d. Want: 2", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"log": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
	"server": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			handler := func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "<h1>Hello World</h1>")
			}

			http.HandleFunc("/", handler)
			log.Fatal(http.ListenAndServe(":8080", nil))
			return NULL
		},
	},
	"file": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. Got: %d. Want: 1", len(args))
			}

			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `file` must be STRING, got %s", args[0].Type())
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

			return &object.String{Value: string(file[:])}
		},
	},
}

func getBuiltin(name string, env *object.Environment) (*object.Builtin, bool) {
	if env.Silent && name == "log" {
		return &object.Builtin{
			Fn: func(args ...object.Object) object.Object {
				return NULL
			},
		}, true
	}

	if builtin, ok := builtins[name]; ok {
		return builtin, true
	}

	return builtins[name], false
}
