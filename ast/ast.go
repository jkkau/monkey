package ast

import (
	"monkey/scanner"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	return "Program"
}

// example: let identifier = expression;
// LetStatement实现Statement接口
type LetStatement struct {
	Token scanner.Token
	Name  Identifier
	Value Expression
}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Lexeme
}

// 该接口用于标识这是一个Statement
func (l *LetStatement) StatementNode() {
}

// 把Identifier也当成一个Expression，便于简单实现
type Identifier struct {
	Token scanner.Token
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Lexeme
}

func (i *Identifier) ExpressionNode() {
}

// example: return [expression]
// ReturnStatement实现Statement接口
type ReturnStatement struct {
	Token scanner.Token
	Value Expression
}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Lexeme
}

func (r *ReturnStatement) StatementNode() {
}
