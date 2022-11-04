package lexer

import (
	"testing"

	"github.com/chau-t-tran/bengo/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LexerTestSuite struct {
	suite.Suite
}

func assertLexerEquals(T *testing.T, expected []token.Token, lexer *Lexer) {
	/*
		Asserts that the lexer generator produces the correct
		token output and that it terminates.
	*/
	for i, _ := range expected {
		t, err := lexer.NextToken()
		assert.NoError(T, err)
		assert.Equal(T, expected[i], t)
	}
	_, err := lexer.NextToken()
	assert.Error(T, err)
}

func (suite *LexerTestSuite) TestConstructor() {
	mock := NewLexer("i123e")
	expected := Lexer{
		index:      0,
		state:      0,
		byteLength: 0,
		chars:      []rune("i123e"),
	}
	assert.IsType(suite.T(), expected, mock)
	assert.Equal(suite.T(), expected.chars, mock.chars)
}

func (suite *LexerTestSuite) TestParseDigits() {
	lexer := NewLexer("1324e")
	assert.Equal(suite.T(), lexer.parseDigits(), "1324")
}

func (suite *LexerTestSuite) TestParseBytes() {
	lexer := NewLexer("5:hello")
	lexer.index = 2
	lexer.byteLength = 5
	lexer.state = expectingBytes
	assert.Equal(suite.T(), lexer.parseBytes(), "hello")
}

func (suite *LexerTestSuite) TestByte() {
	lexer := NewLexer("4:spam")
	expected := []token.Token{
		{Type: token.BYTE_LENGTH, Literal: "4"},
		{Type: token.COLON, Literal: ":"},
		{Type: token.BYTE_CONTENT, Literal: "spam"},
	}
	assertLexerEquals(suite.T(), expected, &lexer)
}

func (suite *LexerTestSuite) TestInteger() {
	lexer := NewLexer("i123e")
	expected := []token.Token{
		{Type: token.INT_ENTRY, Literal: "i"},
		{Type: token.INT_VALUE, Literal: "123"},
		{Type: token.END, Literal: "e"},
	}
	assertLexerEquals(suite.T(), expected, &lexer)
}

func (suite *LexerTestSuite) TestList() {
	lexer := NewLexer("l4:spami42ee")
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
	assertLexerEquals(suite.T(), expected, &lexer)
}

func (suite *LexerTestSuite) TestDict() {
	lexer := NewLexer("d3:bar4:spam3:fooi42ee")
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
	assertLexerEquals(suite.T(), expected, &lexer)
}

func TestLexerTestSuite(t *testing.T) {
	suite.Run(t, new(LexerTestSuite))
}
