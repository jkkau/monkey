package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	token := NewToken(TokenTypeEOF, "", 0)
	assert.Equal(t, token.Type, TokenTypeEOF)
	assert.Equal(t, token.Lexeme, "")
	assert.Equal(t, token.Line, uint32(0))
}
