package parser

import (
	"monkey/ast"
	"monkey/scanner"
	"testing"

	"github.com/stretchr/testify/assert"
)

type expectedValues struct {
	tokenLitenral string
	identifier    string
	value         string
}

func testLetStatement(t *testing.T, s ast.Statement, expects expectedValues) {
	// assert.Equal(t, "let", s.TokenLiteral())

	letStatement, ok := s.(*ast.LetStatement)
	assert.Equal(t, true, ok)
	assert.Equal(t, letStatement.TokenLiteral(), expects.tokenLitenral)
	assert.Equal(t, letStatement.Name.TokenLiteral(), expects.identifier)
}

func TestLetStatement(t *testing.T) {
	input := `
	let a = 5;
	let b = 10;
	let foobar = 10056;
	`

	scanner := scanner.NewScanner(input)
	tokens := scanner.ScanTokens()
	parser := NewParser(tokens)
	program := parser.ParseProgram()
	assert.Equal(t, 3, len(program.Statements))

	expectLiteral := []expectedValues{
		{tokenLitenral: "let", identifier: "a"},
		{tokenLitenral: "let", identifier: "b"},
		{tokenLitenral: "let", identifier: "foobar"},
	}

	for i := 0; i < len(program.Statements); i++ {
		testLetStatement(t, program.Statements[i], expectLiteral[i])
	}
}

func testReturnStatement(t *testing.T, s ast.Statement) {
	returnStatement, ok := s.(*ast.ReturnStatement)
	assert.Equal(t, true, ok)
	assert.Equal(t, returnStatement.TokenLiteral(), "return")
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 10 5;
	return;
	`

	scanner := scanner.NewScanner(input)
	tokens := scanner.ScanTokens()
	parser := NewParser(tokens)
	program := parser.ParseProgram()
	assert.Equal(t, 4, len(program.Statements))

	for i := 0; i < len(program.Statements); i++ {
		testReturnStatement(t, program.Statements[i])
	}
}
