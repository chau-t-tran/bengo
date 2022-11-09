package token

type Token struct {
	Type    int
	Literal string
}

const (
	NULL int = iota
	TERMINATE
	BYTE_LENGTH
	BYTE_CONTENT
	INT_ENTRY
	INT_VALUE
	LIST_ENTRY
	DICT_ENTRY
	COLON
	END
)

func NewToken(tokenType int, tokenLiteral string) Token {
	return Token{
		Type:    tokenType,
		Literal: tokenLiteral,
	}
}
