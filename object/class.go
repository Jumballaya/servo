package object

import (
	"fmt"

	"github.com/jumballaya/servo/ast"
)

type ClassMethod interface {
	Object
	ClassMethod()
}

type Class struct {
	Name    string
	Parent  *Class
	Fields  []*ast.LetStatement
	Methods map[string]ClassMethod
}

func (c *Class) Inspect() string  { return "class " + c.Name }
func (c *Class) Type() ObjectType { return CLASS_OBJ }
func (c *Class) GetMethod(name string) ClassMethod {
	m, ok := c.Methods[name]

	if ok || c.Parent == nil {
		return m
	}
	return c.Parent.GetMethod(name)
}

type Instance struct {
	Class  *Class
	Fields *Environment
}

func (i *Instance) Inspect() string                   { return fmt.Sprintf("instance of %s", i.Class.Name) }
func (i *Instance) Type() ObjectType                  { return INSTANCE_OBJ }
func (i *Instance) GetMethod(name string) ClassMethod { return i.Class.GetMethod(name) }

// InstanceOf checks to see if the instance is an instance of the given class name
func InstanceOf(class string, i *Instance) bool {
	if i == nil {
		return false
	}

	c := i.Class
	for {
		if c.Name == class {
			return true
		}
		if c.Parent == nil {
			return false
		}
		c = c.Parent
	}
}

// InstanceOfAny checks an object to see if it is an instance of any of the class names provided
func InstanceOfAny(instance Object, classes ...string) bool {
	i, ok := instance.(*Instance)
	if !ok {
		return false
	}

	for _, class := range classes {
		if InstanceOf(class, i) {
			return true
		}
	}

	return false
}

// GetSelf returns the instance belonging to the name 'self'
func GetSelf(env *Environment) *Instance {
	self, ok := env.Get("this")
	if !ok {
		return nil
	}

	i, ok := self.(*Instance)
	if !ok {
		return nil
	}

	return i
}
