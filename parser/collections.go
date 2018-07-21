package parser

import (
	"fmt"

	"github.com/jumballaya/servo/ast"
	"github.com/jumballaya/servo/token"
)

// Parse Array Literal builds each expression for each array element
func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.curToken}
	array.Elements = p.parseExpressionList(token.RBRACKET)
	return array
}

// Parse Hash Literal builds the expressions for the hash pairs and
// packages it into the ast.HashLiteral
func (p *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{Token: p.curToken}
	hash.Pairs = make(map[ast.Expression]ast.Expression)

	for !p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		key := p.parseExpression(LOWEST)

		if !p.expectPeek(token.COLON) {
			msg := fmt.Sprintf("hash declaration missing ':'")
			p.errors = append(p.errors, msg)
			return nil
		}

		p.nextToken()
		value := p.parseExpression(LOWEST)
		hash.Pairs[key] = value

		if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
			msg := fmt.Sprintf("hash declaration must include ',' or '}'")
			p.errors = append(p.errors, msg)
			return nil
		}

	}

	if !p.expectPeek(token.RBRACE) {
		msg := fmt.Sprintf("hash declaration must end with '}'")
		p.errors = append(p.errors, msg)
		return nil
	}

	return hash
}

// Parse Index Expression builds the expression that will be evaluated into the desired index
// example: `arr[1 + 3]` this would create the (1 + 3) expression to be evaluated into (4)
func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	exp := &ast.IndexExpression{Token: p.curToken, Left: left}

	p.nextToken()
	exp.Index = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RBRACKET) {
		msg := fmt.Sprintf("index expressions must end with ']'")
		p.errors = append(p.errors, msg)
		return nil
	}

	return exp
}

// Parse Attribute Expression
func (p *Parser) parseAttributeExpression(left ast.Expression) ast.Expression {
	exp := &ast.AttributeExpression{Token: p.curToken, Left: left}
	p.nextToken()
	i := p.parseExpression(ASSIGN)

	ident, ok := i.(*ast.Identifier)
	if !ok {
		msg := fmt.Sprintf("attribute operator requires an identifier")
		p.errors = append(p.errors, msg)
		return nil
	}

	t := token.Token{Type: token.STRING, Literal: ident.Value}
	exp.Index = &ast.StringLiteral{Token: t, Value: ident.Value}

	return exp
}
