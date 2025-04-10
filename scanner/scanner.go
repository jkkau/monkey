package scanner

import (
	"errors"
	"fmt"
)

type Scanner struct {
	source string
	pos    int // ScanTokens时记录当前处理的字符所在位置
}

func NewScanner(inputSourceCode string) *Scanner {
	return &Scanner{
		source: inputSourceCode,
		pos:    0,
	}
}

func (s *Scanner) next() byte {
	c := s.source[s.pos]
	s.pos++
	return c
}

func (s *Scanner) peek() byte {
	if s.pos < len(s.source) {
		return s.source[s.pos]
	}
	return 0
}

func (s *Scanner) ignoreComments() {
	for s.pos < len(s.source) && s.source[s.pos] != '\n' {
		s.pos++
	}
}

func (s *Scanner) handleString(tokens []Token, line uint32) ([]Token, error) {
	leftPos := s.pos
	for s.source[s.pos] != '"' {
		if s.pos >= len(s.source) {
			return nil, errors.New("no right \"")
		}
		s.pos++
	}
	str := s.source[leftPos:s.pos]
	tokens = append(tokens, NewToken(TokenTypeString, str, line))

	// pos移到'"'后面，以处理后面的字符
	s.pos++
	return tokens, nil
}

func (s *Scanner) getIdentifier() string {
	leftPos := s.pos
	for s.pos < len(s.source) && isAlphaNumberOrUnderscore(s.source[s.pos]) {
		s.pos++
	}
	return s.source[leftPos-1 : s.pos]
}

func (s *Scanner) handleIdentifier(tokens []Token, line uint32) []Token {
	str := s.getIdentifier()
	if str == "let" {
		tokens = append(tokens, NewToken(TokenTypeLet, "let", line))
	} else if str == "fn" {
		tokens = append(tokens, NewToken(TokenTypeFunc, "fn", line))
	} else if str == "and" {
		tokens = append(tokens, NewToken(TokenTypeAnd, "and", line))
	} else if str == "or" {
		tokens = append(tokens, NewToken(TokenTypeOr, "or", line))
	} else if str == "true" {
		tokens = append(tokens, NewToken(TokenTypeTrue, "true", line))
	} else if str == "false" {
		tokens = append(tokens, NewToken(TokenTypeFalse, "false", line))
	} else if str == "return" {
		tokens = append(tokens, NewToken(TokenTypeReturn, "return", line))
	} else if str == "if" {
		tokens = append(tokens, NewToken(TokenTypeIf, "if", line))
	} else if str == "else" {
		tokens = append(tokens, NewToken(TokenTypeElse, "else", line))
	} else if str == "while" {
		tokens = append(tokens, NewToken(TokenTypeWhile, "while", line))
	} else if str == "for" {
		tokens = append(tokens, NewToken(TokenTypeFor, "for", line))
	} else if isNumber(str) {
		tokens = append(tokens, NewToken(TokenTypeNumber, str, line))
	} else {
		tokens = append(tokens, NewToken(TokenTypeIdentifier, str, line))
	}
	return tokens
}

func (s *Scanner) ScanTokens() []Token {
	var tokens []Token
	line := uint32(1)
	for s.pos < len(s.source) {
		c := s.next()
		switch c {
		case '+':
			tokens = append(tokens, NewToken(TokenTypePlus, "+", line))
		case '-':
			tokens = append(tokens, NewToken(TokenTypeMinus, "-", line))
		case '*':
			tokens = append(tokens, NewToken(TokenTypeStar, "*", line))
		case '{':
			tokens = append(tokens, NewToken(TokenTypeLBrace, "{", line))
		case '}':
			tokens = append(tokens, NewToken(TokenTypeRBrace, "}", line))
		case '(':
			tokens = append(tokens, NewToken(TokenTypeLParen, "(", line))
		case ')':
			tokens = append(tokens, NewToken(TokenTypeRParen, ")", line))
		case ';':
			tokens = append(tokens, NewToken(TokenTypeSemicolon, ";", line))
		case ',':
			tokens = append(tokens, NewToken(TokenTypeComma, ",", line))
		case '\n':
			line++
		case '=':
			if s.peek() == '=' {
				s.pos++
				tokens = append(tokens, NewToken(TokenTypeEqualEqual, "==", line))
			} else {
				tokens = append(tokens, NewToken(TokenTypeEqual, "=", line))
			}
		case '!':
			if s.peek() == '=' {
				s.pos++
				tokens = append(tokens, NewToken(TokenTypeBangEqual, "!=", line))
			} else {
				tokens = append(tokens, NewToken(TokenTypeBang, "!", line))
			}
		case '<':
			if s.peek() == '=' {
				s.pos++
				tokens = append(tokens, NewToken(TokenTypeLessEqual, "<=", line))
			} else {
				tokens = append(tokens, NewToken(TokenTypeLess, "<", line))
			}
		case '>':
			if s.peek() == '=' {
				s.pos++
				tokens = append(tokens, NewToken(TokenTypeGreaterEqual, ">=", line))
			} else {
				tokens = append(tokens, NewToken(TokenTypeGreater, ">", line))
			}
		case '/':
			if s.peek() == '/' {
				s.pos++
				s.ignoreComments()
			} else {
				tokens = append(tokens, NewToken(TokenTypeSlash, "/", line))
			}
		case '"':
			/*
				写成if tokens, err := s.handleString(tokens, line); err != nil {}会有问题
				这么写会把tokens认为是一个只在if作用域生效的临时变量
			*/
			var err error
			if tokens, err = s.handleString(tokens, line); err != nil {
				fmt.Printf("line: %d, error: %s\n", line, err.Error())
				return tokens
			}
		default:
			if isAlphaNumberOrUnderscore(c) {
				tokens = s.handleIdentifier(tokens, line)
			}
		}
	}
	tokens = append(tokens, NewToken(TokenTypeEOF, "", line))
	return tokens
}

func isAlphaNumberOrUnderscore(c byte) bool {
	if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_' {
		return true
	}
	return false
}

func isNumber(str string) bool {
	for _, c := range str {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
