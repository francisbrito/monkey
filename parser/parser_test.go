package parser

import (
	"monkey/ast"
	"monkey/lexer"
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
