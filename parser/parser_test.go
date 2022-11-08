package parser

import (
	"testing"

	"github.com/chau-t-tran/bengo/ast"
	"github.com/chau-t-tran/bengo/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ParserTestSuite struct {
	suite.Suite
}

func (suite *ParserTestSuite) TestNextToken() {
	lexer := NewLexer("4:spam")
	parser := newParser(lexer)

	expectedLength, err := token.NewToken(token.BYTE_LENGTH, "4")
	assert.NoError(suite.T(), err)
	expectedColon, err := token.NewToken(token.COLON, ":")
	assert.NoError(suite.T(), err)
	expectedBytes, err := token.NewToken(BYTE_LENGTH, "spam")
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), parser.NextToken(), expectedLength)
	assert.Equal(suite.T(), parser.index, 1)
	assert.Equal(suite.T(), parser.NextToken(), expectedColon)
	assert.Equal(suite.T(), parser.index, 2)
	assert.Equal(suite.T(), parser.NextToken(), expectedBytes)
	assert.Equal(suite.T(), parser.index, 6)
}

func (suite *ParserTestSuite) TestPeekToken() {
	lexer := NewLexer("4:spam")
	parser := newParser(lexer)

	expectedLength, err := token.NewToken(token.BYTE_LENGTH, "4")
	assert.NoError(suite.T(), err)
	expectedColon, err := token.NewToken(token.COLON, ":")
	assert.NoError(suite.T(), err)
	expectedBytes, err := token.NewToken(BYTE_LENGTH, "spam")
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), parser.PeekToken(), expectedLength)
	assert.Equal(suite.T(), parser.index, 0)
	parser.NextToken()
	assert.Equal(suite.T(), parser.PeekToken(), expectedColon)
	assert.Equal(suite.T(), parser.index, 1)
	parser.NextToken()
	assert.Equal(suite.T(), parser.PeekToken(), expectedBytes)
	assert.Equal(suite.T(), parser.index, 2)
}

func (suite *ParserTestSuite) TestParseByte() {
	lexer := NewLexer("4:spam")
	parser := newParser(lexer)
	actual := parser.parseByte()
	expected := ast.NewByteNode("spam")
	assert.Equal(suite.T(), expected, actual)
}

func TestParserTestSuite(t *testing.T) {
	suite.Run(t, new(ParserTestSuite))
}
