package token

type Token struct {
	Type    int
	Literal string
}

const (
	BYTE_LENGTH  = 0
	BYTE_CONTENT = 1
	INT_ENTRY    = 2
	INT_VALUE    = 3
	LIST_ENTRY   = 4
	DICT_ENTRY   = 5
	COLON        = 6
	END          = 7
)

func NewToken(tokenType int, tokenLiteral string) Token {
	return Token{
		Type:    tokenType,
		Literal: tokenLiteral,
	}
}
