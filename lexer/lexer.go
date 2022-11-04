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
		err = errors.New("End of input")
		return
	}

	c := l.chars[l.index]
	switch c {
	case 'i':
		t = token.NewToken(token.INT_ENTRY, string(c))
		l.state = expectingInt
		l.index += 1
	case 'l':
		t = token.NewToken(token.LIST_ENTRY, string(c))
		l.state = expectingNothing
		l.index += 1
	case 'd':
		t = token.NewToken(token.DICT_ENTRY, string(c))
		l.state = expectingNothing
		l.index += 1
	case 'e':
		t = token.NewToken(token.END, string(c))
		l.state = expectingNothing
		l.index += 1
	case ':':
		t = token.NewToken(token.COLON, string(c))
		l.index += 1
	}

	if t.Literal != "" {
		return
	}

	switch l.state {
	case expectingBytes:
		bytes := l.nextBytes()
		l.state = expectingNothing
		t = token.NewToken(token.BYTE_CONTENT, bytes)
		break
	case expectingInt:
		if unicode.IsDigit(c) {
			digits := l.nextDigits()
			l.state = expectingNothing
			t = token.NewToken(token.INT_VALUE, digits)
		} else {
			err = errors.New("Expected digit after int")
		}
		break
	default:
		if unicode.IsDigit(c) {
			digits := l.nextDigits()
			byteLength, _ := strconv.Atoi(digits)
			l.byteLength = byteLength
			l.state = expectingBytes
			t = token.NewToken(token.BYTE_LENGTH, digits)
		} else {
			err = errors.New("Unexpected character")
		}
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
