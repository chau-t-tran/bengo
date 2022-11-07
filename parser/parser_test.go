package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ParserTestSuite struct {
	suite.Suite
}

func (suite *ParserTestSuite) TestParseByte() {
	lexer := NewLexer("4:spam")
	parser := newParser(lexer)
	actual := parser.parseByte()
	expected := NewByteNode("spam")
	assert.Equal(suite.T(), expected, actual)
}

func TestParserTestSuite(t *testing.T) {
	suite.Run(t, new(ParserTestSuite))
}
