package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	lexer = LexerConstructor()
)

type LexerTestSuite struct {
	suite.Suite
}

func (suite *LexerTestSuite) TestByte() {
	input := "4:spam"
	expected := []token.token.Token{
		{Type: token.BYTE_LENGTH, Literal: "4"},
		{Type: token.BYTE_CONTENT, Literal: "spam"},
	}
	actual := lexer.Lex(input)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *LexerTestSuite) TestInteger() {
	input := "i123e"
	expected := []token.token.Token{
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
		{Type: token.DICT_ENTRY, Literal: "l"},
		{Type: token.BYTE_LENGTH, Literal: "3"},
		{Type: token.BYTE_CONTENT, Literal: "bar"},
		{Type: token.BYTE_LENGTH, Literal: "4"},
		{Type: token.BYTE_CONTENT, Literal: "spam"},
		{Type: token.BYTE_LENGTH, Literal: "3"},
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
