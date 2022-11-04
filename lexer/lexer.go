package lexer

import (
	"errors"
	"strconv"
	"unicode"

	"github.com/chau-t-tran/bengo/token"
)

const (
	expectingNothing = 0
	expectingBytes   = 1
	expectingInt     = 2
)

type Lexer struct {
	index      int
	state      int
	byteLength int
	chars      []rune
}

func NewLexer(input string) Lexer {
	return Lexer{
		index:      0,
		state:      0,
		byteLength: 0,
		chars:      []rune(input),
	}
}

func (l *Lexer) NextToken() (t token.Token, err error) {
	if l.index >= len(l.chars) {
		return t, errors.New("End of input")
	}

	c := l.chars[l.index]
	switch c {
	case 'i':
		t = token.NewToken(token.INT_ENTRY, string(c))
		l.state = expectingInt
	case 'l':
		t = token.NewToken(token.LIST_ENTRY, string(c))
		l.state = expectingNothing
	case 'd':
		t = token.NewToken(token.DICT_ENTRY, string(c))
		l.state = expectingNothing
	case 'e':
		t = token.NewToken(token.END, string(c))
		l.state = expectingNothing
	case ':':
		t = token.NewToken(token.COLON, string(c))
	}

	if t.Literal != "" {
		l.index += 1
		return
	}

	switch l.state {
	case expectingBytes:
		bytes := l.nextBytes()
		l.state = expectingNothing
		t = token.NewToken(token.BYTE_CONTENT, bytes)
		break
	case expectingInt:
		if !unicode.IsDigit(c) {
			return t, errors.New("Expected digit after int")
		}
		digits := l.nextDigits()
		l.state = expectingNothing
		t = token.NewToken(token.INT_VALUE, digits)
		break
	default:
		if !unicode.IsDigit(c) {
			return t, errors.New("Expected digit in byte length definition")
		}
		digits := l.nextDigits()
		byteLength, atoiErr := strconv.Atoi(digits)
		if atoiErr != nil {
			return t, atoiErr
		}
		l.byteLength = byteLength
		l.state = expectingBytes
		t = token.NewToken(token.BYTE_LENGTH, digits)
	}
	return
}

func (l *Lexer) nextDigits() string {
	digits := []rune{}
	for unicode.IsDigit(l.chars[l.index]) {
		digits = append(digits, l.chars[l.index])
		l.index += 1
	}
	return string(digits)
}

func (l *Lexer) nextBytes() string {
	bytes := []rune{}
	for i := 0; i < l.byteLength; i++ {
		bytes = append(bytes, l.chars[l.index])
		l.index += 1
	}
	return string(bytes)
}
