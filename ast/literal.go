package ast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/jumballaya/servo/token"
)

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FloatLiteral) String() string       { return fl.Token.Literal }

type CommentLiteral struct {
	Token token.Token
	Value string
}

func (cl *CommentLiteral) expressionNode()      {}
func (cl *CommentLiteral) TokenLiteral() string { return cl.Token.Literal }
func (cl *CommentLiteral) String() string       { return cl.Token.Literal }

type ArrayLiteral struct {
	Token    token.Token // '[' token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type HashLiteral struct {
	Token token.Token // '{' token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode()      {}
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer
	pairs := []string{}

	for key, val := range hl.Pairs {
		pairs = append(pairs, key.String()+": "+val.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token // 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (b *BooleanLiteral) expressionNode()      {}
func (b *BooleanLiteral) TokenLiteral() string { return b.Token.Literal }
func (b *BooleanLiteral) String() string       { return b.Token.Literal }

type NullLiteral struct {
	Token token.Token
}

func (n *NullLiteral) expressionNode()      {}
func (n *NullLiteral) TokenLiteral() string { return n.Token.Literal }
func (n *NullLiteral) String() string       { return n.Token.Literal }

type ClassLiteral struct {
	Token   token.Token
	Name    string
	Parent  string
	Fields  []*LetStatement
	Methods map[string]*FunctionLiteral
}

func (c *ClassLiteral) expressionNode()      {}
func (c *ClassLiteral) TokenLiteral() string { return c.Token.Literal }
func (c *ClassLiteral) String() string {
	return fmt.Sprintf("class %s::%s {...}", c.Name, c.Parent)
}

type InstanceLiteral struct {
	Class     Expression
	Arguments []Expression
}

func (il *InstanceLiteral) expressionNode()      {}
func (il *InstanceLiteral) TokenLiteral() string { return "new" }
func (il *InstanceLiteral) String() string {
	args := []string{}
	for _, a := range il.Arguments {
		args = append(args, a.String())
	}

	return fmt.Sprintf("new %s(%s)", il.Class, strings.Join(args, ", "))
}
