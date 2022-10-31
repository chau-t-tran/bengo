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
	expected := []token.Token{
		{Name: token.BYTE_LENGTH, Literal: "4"},
		{Name: token.BYTE_CONTENT, Literal: "spam"},
	}
	actual := lexer.Lex(input)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *LexerTestSuite) TestInteger() {
	input := "i123e"
	expected := []token.Token{
		{Name: token.INT_ENTRY, Literal: "i"},
		{Name: token.INT_VALUE, Literal: "123"},
		{Name: token.END, Literal: "e"},
	}
	actual := lexer.Lex(input)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *LexerTestSuite) TestList() {
	input := "l4:spami42ee"
	expected := []Token{
		{Name: token.LIST_ENTRY, Literal: "l"},
		{Name: token.BYTE_LENGTH, Literal: "4"},
		{Name: token.BYTE_CONTENT, Literal: "spam"},
		{Name: token.INT_ENTRY, Literal: "i"},
		{Name: token.INT_VALUE, Literal: "42"},
		{Name: token.END, Literal: "e"},
		{Name: token.END, Literal: "e"},
	}
	actual := lexer.Lex(input)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *LexerTestSuite) TestDict() {
	input := "d3:bar4:spam3:fooi42ee"
	expected := []Token{
		{Name: token.DICT_ENTRY, Literal: "l"},
		{Name: token.BYTE_LENGTH, Literal: "3"},
		{Name: token.BYTE_CONTENT, Literal: "bar"},
		{Name: token.BYTE_LENGTH, Literal: "4"},
		{Name: token.BYTE_CONTENT, Literal: "spam"},
		{Name: token.BYTE_LENGTH, Literal: "3"},
		{Name: token.BYTE_CONTENT, Literal: "foo"},
		{Name: token.INT_ENTRY, Literal: "i"},
		{Name: token.INT_VALUE, Literal: "42"},
		{Name: token.END, Literal: "e"},
		{Name: token.END, Literal: "e"},
	}
	actual := lexer.Lex(input)
	assert.Equal(suite.T(), expected, actual)
}

func TestLexerTestSuite(t *testing.T) {
	suite.Run(t, new(LexerTestSuite))
}
