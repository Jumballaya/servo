package parser

import (
	"fmt"

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
	case token.CLASS:
		return p.parseClassStatement()
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

// Parse Reassign takes a reassignment operator like `[identifier] += value` and turns it into
// a brand new assignment expression
func (p *Parser) parseReassignExpression(left ast.Expression) ast.Expression {
	stmt := &ast.AssignExpression{Token: p.curToken, Left: left}
	p.nextToken()
	right := p.parseExpression(LOWEST).(ast.Expression)

	switch stmt.Token.Type {
	case token.PLUSASSIGN:
		stmt.Value = makeInfix(token.PLUS, left, right)
	case token.MINUSASSIGN:
		stmt.Value = makeInfix(token.MINUS, left, right)
	case token.ASTERISKASSIGN:
		stmt.Value = makeInfix(token.ASTERISK, left, right)
	case token.SLASHASSIGN:
		stmt.Value = makeInfix(token.SLASH, left, right)
	case token.ASSIGN:
		stmt.Value = p.parseExpression(LOWEST).(ast.Expression)
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// Make Infix creates an infix expression
func makeInfix(t token.TokenType, left, right ast.Expression) *ast.InfixExpression {
	return &ast.InfixExpression{
		Token:    token.Token{Type: t, Literal: string(t)},
		Left:     left,
		Operator: string(t),
		Right:    right,
	}
}

// Parse Class Statement
func (p *Parser) parseClassStatement() ast.Statement {
	classToken := p.curToken
	if !p.expectPeek(token.IDENT) {
		msg := fmt.Sprintf("could not parse %q as an identifier", p.peekToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	stmt := &ast.LetStatement{Token: createKeywordToken("let")}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	p.insertToken(classToken)
	p.nextToken()

	exp, ok := p.parseExpression(LOWEST).(ast.Expression)
	if !ok {
		return nil
	}
	stmt.Value = exp

	if class, ok := stmt.Value.(*ast.ClassLiteral); ok {
		class.Name = stmt.Name.String()
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
