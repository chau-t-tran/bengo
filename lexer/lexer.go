package lexer

import (
	"strconv"
	"unicode"

	"github.com/chau-t-tran/bengo/token"
)

type Lexer struct {
	curr   int
	chars  []rune
	tokens []token.Token
}

func NewLexer() Lexer {
	return Lexer{
		curr:   0,
		chars:  []rune{},
		tokens: []token.Token{},
	}
}

func (l *Lexer) Lex(input string) []token.Token {
	l.clear()
	l.chars = []rune(input)
	l.curr = 0
	l.parseUnknown()
	return l.tokens
}

func (l *Lexer) clear() {
	l.curr = 0
	l.chars = []rune{}
	l.tokens = []token.Token{}
}

func (l *Lexer) parseUnknown() error {
	for l.curr < len(l.chars) {
		c := l.chars[l.curr]
		switch c {
		case 'i':
			entry := token.NewToken(token.INT_ENTRY, string(c))
			// expect digits
			l.curr += 1
			digits := l.parseDigits()
			value := token.NewToken(token.INT_VALUE, digits)
			l.tokens = append(l.tokens, entry, value)
		case 'l':
			literal := token.NewToken(token.LIST_ENTRY, string(c))
			l.tokens = append(l.tokens, literal)
		case 'd':
			literal := token.NewToken(token.DICT_ENTRY, string(c))
			l.tokens = append(l.tokens, literal)
		case ':':
			literal := token.NewToken(token.COLON, string(c))
			l.tokens = append(l.tokens, literal)
		case 'e':
			literal := token.NewToken(token.END, string(c))
			l.tokens = append(l.tokens, literal)
		default:
			if unicode.IsDigit(c) {
				lengthString := l.parseDigits()
				lengthToken := token.NewToken(token.BYTE_LENGTH, lengthString)
				length, _ := strconv.Atoi(lengthString)
				l.curr += 1

				colonToken := token.NewToken(token.COLON, string(l.chars[l.curr]))
				l.curr += 1

				bytes := l.parseBytes(length)
				bytesToken := token.NewToken(token.BYTE_CONTENT, bytes)
				l.tokens = append(l.tokens, lengthToken, colonToken, bytesToken)
			}
		}
		l.curr += 1
	}
	return nil
}

func (l *Lexer) parseDigits() string {
	digits := []rune{}
	for unicode.IsDigit(l.chars[l.curr]) {
		digits = append(digits, l.chars[l.curr])
		l.curr += 1
	}
	// place pointer right behind nex
	l.curr -= 1
	return string(digits)
}

func (l *Lexer) parseBytes(byteLength int) string {
	bytes := []rune{}
	for i := 0; i < byteLength; i++ {
		bytes = append(bytes, l.chars[l.curr])
		l.curr += 1
	}
	// place pointer right behind nex
	l.curr -= 1
	return string(bytes)
}
