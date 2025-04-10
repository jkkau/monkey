package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanner(t *testing.T) {
	// str := `/+-*{}();,.=!<>!=>=<===//====*****+-*
	// 	   /"abc"
	// 	   123
	//
	// 	   true
	// 	   fn
	// 	   abc let`
	str := `/+-*{}();,=!<>!=>=<===//====*****+-*
	          /"abc"
				123
			   *
				true
				fn
				abc let`

	expectTokens := []Token{
		NewToken(TokenTypeSlash, "/", 1),
		NewToken(TokenTypePlus, "+", 1),
		NewToken(TokenTypeMinus, "-", 1),
		NewToken(TokenTypeStar, "*", 1),
		NewToken(TokenTypeLBrace, "{", 1),
		NewToken(TokenTypeRBrace, "}", 1),
		NewToken(TokenTypeLParen, "(", 1),
		NewToken(TokenTypeRParen, ")", 1),
		NewToken(TokenTypeSemicolon, ";", 1),
		NewToken(TokenTypeComma, ",", 1),
		NewToken(TokenTypeEqual, "=", 1),
		NewToken(TokenTypeBang, "!", 1),
		NewToken(TokenTypeLess, "<", 1),
		NewToken(TokenTypeGreater, ">", 1),
		NewToken(TokenTypeBangEqual, "!=", 1),
		NewToken(TokenTypeGreaterEqual, ">=", 1),
		NewToken(TokenTypeLessEqual, "<=", 1),
		NewToken(TokenTypeEqualEqual, "==", 1),
		NewToken(TokenTypeSlash, "/", 2),
		NewToken(TokenTypeString, "abc", 2),
		NewToken(TokenTypeNumber, "123", 3),
		NewToken(TokenTypeStar, "*", 4),
		NewToken(TokenTypeTrue, "true", 5),
		NewToken(TokenTypeFunc, "fn", 6),
		NewToken(TokenTypeIdentifier, "abc", 7),
		NewToken(TokenTypeLet, "let", 7),
		NewToken(TokenTypeEOF, "", 7),
	}

	scanner := NewScanner(str)
	tokens := scanner.ScanTokens()
	assert.Equal(t, expectTokens, tokens)
}

func TestIsNumber(t *testing.T) {
	assert.True(t, isNumber("123"))
}
