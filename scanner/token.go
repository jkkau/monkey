package scanner

import "fmt"

type TokenType int

const (
	TokenTypeEOF          TokenType = iota
	TokenTypePlus                   // +
	TokenTypeMinus                  // -
	TokenTypeStar                   // *
	TokenTypeSlash                  // /
	TokenTypeLBrace                 // {
	TokenTypeRBrace                 // }
	TokenTypeLParen                 // (
	TokenTypeRParen                 // )
	TokenTypeSemicolon              // ;
	TokenTypeComma                  // ,
	TokenTypeBang                   // !
	TokenTypeBangEqual              // !=
	TokenTypeLess                   // <
	TokenTypeLessEqual              // <=
	TokenTypeGreater                // >
	TokenTypeGreaterEqual           // >=
	TokenTypeEqual                  // =
	TokenTypeEqualEqual             // ==
	TokenTypeString                 // "xxx"
	TokenTypeNumber                 // 123
	TokenTypeIdentifier             // identifier
	TokenTypeIf                     // if
	TokenTypeElse                   // else
	TokenTypeWhile                  // while
	TokenTypeFor                    // for
	TokenTypeReturn                 // return
	TokenTypeAnd                    // and
	TokenTypeOr                     // or
	TokenTypeTrue                   // true
	TokenTypeFalse                  // false
	TokenTypeFunc                   // fn
	TokenTypeLet                    // let
)

var tokenMap = map[TokenType]string{
	TokenTypeEOF:          "EOF",
	TokenTypePlus:         "PLUS",
	TokenTypeMinus:        "MINUS",
	TokenTypeStar:         "STAR",
	TokenTypeSlash:        "SLASH",
	TokenTypeLBrace:       "LBRACE",
	TokenTypeRBrace:       "RBRACE",
	TokenTypeLParen:       "LPAREN",
	TokenTypeRParen:       "RPAREN",
	TokenTypeSemicolon:    "SEMICOLON",
	TokenTypeComma:        "COMMA",
	TokenTypeBang:         "BANG",
	TokenTypeBangEqual:    "BANGEQUAL",
	TokenTypeLess:         "LESS",
	TokenTypeLessEqual:    "LESSEQUAL",
	TokenTypeGreater:      "GREATER",
	TokenTypeGreaterEqual: "GREATEREQUAL",
	TokenTypeEqual:        "EQUAL",
	TokenTypeEqualEqual:   "EQUALEQUAL",
	TokenTypeString:       "STRING",
	TokenTypeNumber:       "NUMBER",
	TokenTypeIdentifier:   "IDENTIFIER",
	TokenTypeFunc:         "FUNC",
	TokenTypeReturn:       "RETURN",
	TokenTypeTrue:         "TRUE",
	TokenTypeFalse:        "FALSE",
	TokenTypeLet:          "LET",
}

type Token struct {
	Type   TokenType
	Lexeme string
	Line   uint32
}

func NewToken(t TokenType, s string, l uint32) Token {
	return Token{
		Type:   t,
		Lexeme: s,
		Line:   l,
	}
}

func (t *Token) String() string {
	return "{" + tokenMap[t.Type] + "," + t.Lexeme + "," + fmt.Sprint(t.Line) + "}"
}
