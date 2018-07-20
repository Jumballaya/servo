package parser

import (
	"github.com/jumballaya/servo/ast"
	"github.com/jumballaya/servo/token"
)

// Parse Statement checks the statement type and runs the corresponding parsing function
// Currently only 3 types of statements exist:
//		1. Let statements
//    2. Return statements
//    3. Expression statements (everything else)
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// Parse Let Statement attempts to assign a right-hand expression to a left-hand
// identifier using the assignment operator, '='
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// Parse Return Statement attempts to build an expression for the return value
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	stmt.ReturnValue = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// Parse Import Statement attempts to build the import statement
func (p *Parser) parseImportStatement() ast.Expression {
	// token.IMPORT with value 'import'
	stmt := &ast.ImportExpression{Token: p.curToken, Value: p.curToken.Literal}

	// Check for the identifier e.g. map in `import map from './test.svo';`
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// Set the name of the import
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Make sure the next token is a FROM token
	if !p.expectPeek(token.FROM) {
		return nil
	}

	// Make sure the next token is a string
	if !p.expectPeek(token.STRING) {
		return nil
	}

	// Make sure the string isn't empty
	if p.curToken.Literal == "" {
		return nil
	}

	// Set the path from the string
	stmt.Path = &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}

	// Check for SEMICOLON
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
