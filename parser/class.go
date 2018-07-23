package parser

import (
	"github.com/jumballaya/servo/ast"
	"github.com/jumballaya/servo/token"
)

// Parse Class Literal
func (p *Parser) parseClassLiteral() ast.Expression {
	c := &ast.ClassLiteral{
		Fields:  make([]*ast.LetStatement, 0),
		Methods: make(map[string]*ast.FunctionLiteral),
	}

	if p.peekTokenIs(token.COLONCOLON) {
		p.nextToken()
		if !p.expectPeek(token.IDENT) {
			return nil
		}
		c.Parent = p.curToken.Literal
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	body := p.parseBlockStatement()

	for _, stmt := range body.Statements {
		let, ok := stmt.(*ast.LetStatement)
		if !ok {
			return nil
		}

		switch s := let.Value.(type) {
		case *ast.FunctionLiteral:
			c.Methods[let.Name.Value] = s
			c.Fields = append(c.Fields, let)
		default:
			c.Fields = append(c.Fields, let)
		}
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return c
}

// Parse New Expression
func (p *Parser) parseNewExpression() ast.Expression {
	i := &ast.NewInstance{}

	p.nextToken()
	classExp := p.parseExpression(LOWEST)

	call, ok := classExp.(*ast.CallExpression)
	if !ok {
		msg := "invalid class instance creation"
		p.errors = append(p.errors, msg)
		return nil
	}

	i.Class = call.Function
	i.Arguments = call.Arguments

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return i
}
