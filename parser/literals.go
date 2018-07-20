package parser

import (
	"fmt"
	"strconv"

	"github.com/jumballaya/servo/ast"
	"github.com/jumballaya/servo/token"
)

// Parse Integer Literal
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

// Parse String Literal
func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

// Parse Boolean Literal
func (p *Parser) parseBooleanLiteral() ast.Expression {
	return &ast.Boolean{
		Token: p.curToken,
		Value: p.curTokenIs(token.TRUE),
	}
}

// Parse Identifier
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// Parse Null Literal
func (p *Parser) parseNullLiteral() ast.Expression {
	return &ast.NullLiteral{Token: p.curToken}
}

// Parse Comment Literal
func (p *Parser) parseCommentLiteral() ast.Expression {
	return &ast.CommentLiteral{Token: p.curToken, Value: p.curToken.Literal}
}
