package parser

import (
	"monkey/ast"
	"monkey/scanner"
	"testing"

	"github.com/stretchr/testify/assert"
)

type expectedValues struct {
	identifier string
	value      string
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
		{identifier: "a"},
		{identifier: "b"},
		{identifier: "foobar"},
	}

	for i := 0; i < len(program.Statements); i++ {
		testLetStatement(t, program.Statements[i], expectLiteral[i])
	}
}

func testLetStatement(t *testing.T, s ast.Statement, expects expectedValues) {
	// assert.Equal(t, "let", s.TokenLiteral())

	letStatement, ok := s.(*ast.LetStatement)
	assert.Equal(t, true, ok)
	assert.Equal(t, letStatement.Name.TokenLiteral(), expects.identifier)
}
