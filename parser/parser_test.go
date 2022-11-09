package parser

import (
	"testing"

	"github.com/chau-t-tran/bengo/ast"
	"github.com/chau-t-tran/bengo/lexer"
	"github.com/chau-t-tran/bengo/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ParserTestSuite struct {
	suite.Suite
}

func (suite *ParserTestSuite) TestNextToken() {
	lexer := lexer.NewLexer("4:spam")
	parser, err := newParser(lexer)
	assert.NoError(suite.T(), err)

	expectedLength := token.NewToken(token.BYTE_LENGTH, "4")
	expectedColon := token.NewToken(token.COLON, ":")
	expectedBytes := token.NewToken(token.BYTE_LENGTH, "spam")

	actualLength, err := parser.NextToken()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedLength, actualLength)
	assert.Equal(suite.T(), parser.index, 1)

	actualColon, err := parser.NextToken()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedColon, actualColon)
	assert.Equal(suite.T(), parser.index, 2)

	actualBytes, err := parser.NextToken()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedBytes, actualBytes)
	assert.Equal(suite.T(), parser.index, 6)
}

func (suite *ParserTestSuite) TestPeekToken() {
	lexer := lexer.NewLexer("4:spam")
	parser, err := newParser(lexer)
	assert.NoError(suite.T(), err)

	expectedLength := token.NewToken(token.BYTE_LENGTH, "4")
	expectedColon := token.NewToken(token.COLON, ":")
	expectedBytes := token.NewToken(token.BYTE_LENGTH, "spam")

	actualLength := parser.PeekToken()
	assert.Equal(suite.T(), expectedLength, actualLength)
	assert.Equal(suite.T(), parser.index, 0)
	parser.NextToken()

	actualColon := parser.PeekToken()
	assert.Equal(suite.T(), expectedColon, actualColon)
	assert.Equal(suite.T(), parser.index, 1)
	parser.NextToken()

	actualBytes := parser.PeekToken()
	assert.Equal(suite.T(), expectedBytes, actualBytes)
	assert.Equal(suite.T(), parser.index, 2)
}

func (suite *ParserTestSuite) TestParseByte() {
	lexer := lexer.NewLexer("4:spam")
	parser, err := newParser(lexer)
	assert.NoError(suite.T(), err)

	actual, err := parser.parseByte()
	assert.NoError(suite.T(), err)
	expected := ast.NewByteNode("spam")
	assert.Equal(suite.T(), expected, actual)
}

func TestParserTestSuite(t *testing.T) {
	suite.Run(t, new(ParserTestSuite))
}
