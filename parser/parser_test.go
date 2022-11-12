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
	expectedBytes := token.NewToken(token.BYTE_CONTENT, "spam")

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
	expectedBytes := token.NewToken(token.BYTE_CONTENT, "spam")

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

func (suite *ParserTestSuite) TestParseInteger() {
	lexer := lexer.NewLexer("i123456e")
	parser, err := newParser(lexer)
	assert.NoError(suite.T(), err)

	actual, err := parser.parseInt()
	assert.NoError(suite.T(), err)
	expected := ast.NewIntNode("123456")
	assert.Equal(suite.T(), expected, actual)
}

func (suite *ParserTestSuite) TestBasicList() {
	lexer := lexer.NewLexer("li123e4:spam5:helloi23ee")
	parser, err := newParser(lexer)
	assert.NoError(suite.T(), err)

	expected := ast.NewListNode()
	i1 := ast.NewIntNode("123")
	b1 := ast.NewByteNode("spam")
	b2 := ast.NewByteNode("hello")
	i2 := ast.NewIntNode("23")
	expected.Add(i1)
	expected.Add(b1)
	expected.Add(b2)
	expected.Add(i2)

	actual, err := parser.parseList()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *ParserTestSuite) TestNestedList() {
	lexer := lexer.NewLexer("lli32ei33eeli11e4:spamee")
	parser, err := newParser(lexer)
	assert.NoError(suite.T(), err)

	expected := ast.NewListNode()
	l1 := ast.NewListNode()
	i1 := ast.NewIntNode("32")
	i2 := ast.NewIntNode("33")
	l1.Add(i1)
	l1.Add(i2)
	l2 := ast.NewListNode()
	i3 := ast.NewIntNode("11")
	b1 := ast.NewByteNode("spam")
	l2.Add(i3)
	l2.Add(b1)
	expected.Add(l1)
	expected.Add(l2)

	actual, err := parser.parseList()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *ParserTestSuite) TestParseBasicDict() {
	lexer := lexer.NewLexer("d4:val1i32e4:val24:spam6:value3i12ee")
	parser, err := newParser(lexer)
	assert.NoError(suite.T(), err)

	expected := ast.NewDictNode()
	k1 := ast.NewByteNode("val1")
	i1 := ast.NewIntNode("32")
	k2 := ast.NewByteNode("val2")
	b1 := ast.NewByteNode("spam")
	k3 := ast.NewByteNode("value3")
	i2 := ast.NewIntNode("12")
	expected.Add(k1, i1)
	expected.Add(k2, b1)
	expected.Add(k3, i2)

	actual, err := parser.parseDict()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *ParserTestSuite) TestParseUnknownByte() {
	lexer := lexer.NewLexer("4:spam")
	parser, err := newParser(lexer)
	assert.NoError(suite.T(), err)

	actual, err := parser.parseUnknown()
	assert.NoError(suite.T(), err)
	expected := ast.NewByteNode("spam")
	assert.Equal(suite.T(), expected, actual)
}

func (suite *ParserTestSuite) TestParseUnknownInteger() {
	lexer := lexer.NewLexer("i123456e")
	parser, err := newParser(lexer)
	assert.NoError(suite.T(), err)

	actual, err := parser.parseUnknown()
	assert.NoError(suite.T(), err)
	expected := ast.NewIntNode("123456")
	assert.Equal(suite.T(), expected, actual)
}

func (suite *ParserTestSuite) TestParseUnknownList() {
	lexer := lexer.NewLexer("li123e4:spam5:helloi23ee")
	parser, err := newParser(lexer)
	assert.NoError(suite.T(), err)

	expected := ast.NewListNode()
	i1 := ast.NewIntNode("123")
	b1 := ast.NewByteNode("spam")
	b2 := ast.NewByteNode("hello")
	i2 := ast.NewIntNode("23")
	expected.Add(i1)
	expected.Add(b1)
	expected.Add(b2)
	expected.Add(i2)

	actual, err := parser.parseUnknown()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *ParserTestSuite) TestParseUnknownDict() {
	lexer := lexer.NewLexer("d4:val1i32e4:val24:spam6:value3i12ee")
	parser, err := newParser(lexer)
	assert.NoError(suite.T(), err)

	expected := ast.NewDictNode()
	k1 := ast.NewByteNode("val1")
	i1 := ast.NewIntNode("32")
	k2 := ast.NewByteNode("val2")
	b1 := ast.NewByteNode("spam")
	k3 := ast.NewByteNode("value3")
	i2 := ast.NewIntNode("12")
	expected.Add(k1, i1)
	expected.Add(k2, b1)
	expected.Add(k3, i2)

	actual, err := parser.parseUnknown()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func TestParserTestSuite(t *testing.T) {
	suite.Run(t, new(ParserTestSuite))
}
