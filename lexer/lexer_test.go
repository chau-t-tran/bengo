package lexer

import (
	"testing"

	"github.com/chau-t-tran/bengo/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	lexer = NewLexer()
)

type LexerTestSuite struct {
	suite.Suite
}

func (suite *LexerTestSuite) TestConstructor() {
	mock := NewLexer()
	expected := Lexer{
		curr:   0,
		chars:  []rune{},
		tokens: []token.Token{},
	}
	assert.IsType(suite.T(), expected, mock)
}

func (suite *LexerTestSuite) TestClear() {
	mock := Lexer{
		curr:  20,
		chars: []rune{'h'},
		tokens: []token.Token{
			token.Token{
				Type:    token.BYTE_LENGTH,
				Literal: "200",
			},
		},
	}
	mock.clear()
	assert.Equal(suite.T(), 0, mock.curr)
	assert.Equal(suite.T(), 0, len(mock.chars))
	assert.Equal(suite.T(), 0, len(mock.tokens))
}

func (suite *LexerTestSuite) TestParseDigits() {
	lexer.curr = 0
	lexer.chars = []rune("1324e")
	assert.Equal(suite.T(), lexer.parseDigits(), "1324")
}

func (suite *LexerTestSuite) TestParseBytes() {
	lexer.curr = 2
	lexer.chars = []rune("5:hello")
	assert.Equal(suite.T(), lexer.parseBytes(5), "hello")
}

func (suite *LexerTestSuite) TestByte() {
	input := "4:spam"
	expected := []token.Token{
		{Type: token.BYTE_LENGTH, Literal: "4"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.BYTE_CONTENT, Literal: "spam"},
	}
	actual := lexer.Lex(input)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *LexerTestSuite) TestInteger() {
	input := "i123e"
	expected := []token.Token{
		{Type: token.INT_ENTRY, Literal: "i"},
		{Type: token.INT_VALUE, Literal: "123"},
		{Type: token.END, Literal: "e"},
	}
	actual := lexer.Lex(input)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *LexerTestSuite) TestList() {
	input := "l4:spami42ee"
	expected := []token.Token{
		{Type: token.LIST_ENTRY, Literal: "l"},
		{Type: token.BYTE_LENGTH, Literal: "4"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.BYTE_CONTENT, Literal: "spam"},
		{Type: token.INT_ENTRY, Literal: "i"},
		{Type: token.INT_VALUE, Literal: "42"},
		{Type: token.END, Literal: "e"},
		{Type: token.END, Literal: "e"},
	}
	actual := lexer.Lex(input)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *LexerTestSuite) TestDict() {
	input := "d3:bar4:spam3:fooi42ee"
	expected := []token.Token{
		{Type: token.DICT_ENTRY, Literal: "d"},
		{Type: token.BYTE_LENGTH, Literal: "3"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.BYTE_CONTENT, Literal: "bar"},
		{Type: token.BYTE_LENGTH, Literal: "4"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.BYTE_CONTENT, Literal: "spam"},
		{Type: token.BYTE_LENGTH, Literal: "3"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.BYTE_CONTENT, Literal: "foo"},
		{Type: token.INT_ENTRY, Literal: "i"},
		{Type: token.INT_VALUE, Literal: "42"},
		{Type: token.END, Literal: "e"},
		{Type: token.END, Literal: "e"},
	}
	actual := lexer.Lex(input)
	assert.Equal(suite.T(), expected, actual)
}

func TestLexerTestSuite(t *testing.T) {
	suite.Run(t, new(LexerTestSuite))
}
