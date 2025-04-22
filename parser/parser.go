package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/scanner"
)

type Parser struct {
	tokens   []scanner.Token
	curToken *scanner.Token
	pos      uint32
}

func NewParser(ts []scanner.Token) *Parser {
	p := &Parser{
		tokens: ts,
		pos:    0,
	}
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for p.curToken.Type != scanner.TokenTypeEOF {
		st := p.parseStatement()
		if st != nil {
			program.Statements = append(program.Statements, st)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	var st ast.Statement
	switch p.curToken.Type {
	case scanner.TokenTypeLet:
		st = p.parseLetStatement()
	}
	return st
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	if p.curToken.Type != scanner.TokenTypeLet {
		fmt.Printf("token type error: %d\n", p.curToken.Type)
		return nil
	}

	st := &ast.LetStatement{}

	p.nextToken()
	if !p.curTokenIs(scanner.TokenTypeIdentifier) {
		fmt.Printf("parse code error, should be identifier instead of %s\n", p.curToken)
		return nil
	}
	st.Name = ast.Identifier{
		Token: *p.curToken,
	}

	p.nextToken()
	if !p.curTokenIs(scanner.TokenTypeEqual) {
		fmt.Printf("parse code error, should be '=' instead of %s\n", p.curToken)
		return nil
	}

	// TODO: 解析表达式

	p.nextToken()
	for !p.curTokenIs(scanner.TokenTypeSemicolon) {
		p.nextToken()
	}

	return st
}

func (p *Parser) curTokenIs(t scanner.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) nextToken() {
	p.curToken = &p.tokens[p.pos]
	p.pos++
}
