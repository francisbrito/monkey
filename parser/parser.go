package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

const (
	LOWEST      = iota + 1
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)”
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.Type]prefixParseFn
	infixParseFn   map[token.Type]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{errors: []string{}, l: l}
	p.prefixParseFns = map[token.Type]prefixParseFn{}
	p.registerPrefixParseFn(token.IDENT, p.parseIdentifier)
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	prg := &ast.Program{}
	prg.Statements = []ast.Statement{}
	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		prg.Statements = append(prg.Statements, stmt)
		p.nextToken()
	}
	return prg
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

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

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) curTokenIs(t token.Type) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix, ok := p.prefixParseFns[p.curToken.Type]
	if !ok {
		return nil
	}
	le := prefix()
	return le
}

func (p *Parser) peekError(t token.Type) {
	msg := fmt.Sprintf("next token = %q, want = %q", p.peekToken.Type, t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseReturnStatement() ast.Statement {
	rs := &ast.ReturnStatement{Token: p.curToken}
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return rs
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

func (p *Parser) registerPrefixParseFn(t token.Type, f prefixParseFn) {
	p.prefixParseFns[t] = f
}

func (p *Parser) registerInfixParseFn(t token.Type, f infixParseFn) {
	p.infixParseFn[t] = f
}

func (p *Parser) parseExpressionStatement() ast.Statement {
	es := &ast.ExpressionStatement{Token: p.curToken}
	es.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return es
}
