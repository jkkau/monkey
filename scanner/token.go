package scanner

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
