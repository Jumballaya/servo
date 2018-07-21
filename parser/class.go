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
		default:
			c.Fields = append(c.Fields, let)
		}
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return c
}
