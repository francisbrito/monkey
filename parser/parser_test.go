package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 123123;
`
	l := lexer.New(input)
	p := New(l)
	prg := p.ParseProgram()
	if prg == nil {
		t.Fatalf("ParseProgram() = nil")
	}
	checkParserErrors(t, p)
	if sc := len(prg.Statements); sc != 3 {
		t.Fatalf("len(prg.Statements) = %d, want 3", sc)
	}
	tests := []struct {
		ei string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tc := range tests {
		stmt := prg.Statements[i]
		if stmt == nil {
			t.Fatalf("prg.Statements[%d] = nil", i)
		}
		testLetStatement(t, stmt, tc.ei)
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 123123;
`
	l := lexer.New(input)
	p := New(l)
	prg := p.ParseProgram()
	checkParserErrors(t, p)
	if sc := len(prg.Statements); sc != 3 {
		t.Fatalf("len(prg.Statements) = %d, want 3", sc)
	}
	for _, s := range prg.Statements {
		rs, ok := s.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("s = %T, want = *ast.ReturnStatement", s)
		}
		if tl := rs.Token.Literal; tl != "return" {
			t.Errorf(`rs.Token.Literal = %q, want "return"`, tl)
		}
	}
}

func TestString(t *testing.T) {
	prg := &ast.Program{Statements: []ast.Statement{
		&ast.LetStatement{
			Token: token.Token{Literal: "let", Type: token.LET},
			Name: &ast.Identifier{
				Token: token.Token{Literal: "myVar", Type: token.IDENT},
				Value: "myVar",
			},
			Value: &ast.Identifier{
				Token: token.Token{Literal: "anotherVar", Type: token.IDENT},
				Value: "anotherVar",
			},
		},
	}}
	if src := prg.String(); src != "let myVar = anotherVar;" {
		t.Errorf(`prg.String() = %q, want "let myVar = anotherVar;"`, src)
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
	p := New(l)
	prg := p.ParseProgram()
	checkParserErrors(t, p)
	if ls := len(prg.Statements); ls != 1 {
		t.Fatalf("len(prg.Statements) = %d, want 1", ls)
	}
	es, ok := prg.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("prg.Statements[0] = %T, want *ast.ExpressionStatement", es)
	}
	ident, ok := es.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("ident = %T, want *ast.Identifier", ident)
	}
	if ident.Value != "foobar" {
		t.Errorf(`ident.Value = %q, want "foobar"`, ident.Value)
	}
	if tl := ident.TokenLiteral(); tl != "foobar" {
		t.Errorf(`ident.TokenLiteral() = %q, want "foobar"`, tl)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, ei string) {
	if tl := s.TokenLiteral(); tl != "let" {
		t.Errorf(`s.TokenLiteral() = %q, want = "let"`, tl)
	}
	ls, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s = %T, want *ast.LetStatement", s)
	}
	if i := ls.Name.Value; i != ei {
		t.Errorf("ls.Name.Value = %q, want = %q", i, ei)
	}
	if tl := ls.Name.TokenLiteral(); tl != ei {
		t.Errorf("ls.Name.TokenLiteral() = %q, want = %q", tl, ei)
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.errors
	if len(errors) == 0 {
		return
	}
	t.Errorf("len(p.errors) = %d, want 0", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
