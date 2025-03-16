package parser

import (
	"testing"

	"github.com/abhi-abhiram/interpreter/ast"
	"github.com/abhi-abhiram/interpreter/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	
	let a = 5;
	let b = 10;
	let var = 8383;
	
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParserProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"a"},
		{"b"},
		{"var"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, identifier string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", stmt)
		return false
	}

	if letStmt.Name.TokenLiteral() != identifier {
		t.Errorf("s.Name not '%s'. got=%s", identifier, letStmt.Name)
		return false
	}

	return true

}
