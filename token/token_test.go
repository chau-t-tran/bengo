package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TokenTestSuite struct {
	suite.Suite
}

func (suite *TokenTestSuite) TestConstructor() {
	mock := NewToken(BYTE_LENGTH, "2")
	expected := Token{
		Type:    BYTE_LENGTH,
		Literal: "2",
	}
	assert.IsType(suite.T(), expected, mock)
	assert.Equal(suite.T(), expected.Type, mock.Type)
	assert.Equal(suite.T(), expected.Literal, mock.Literal)
}

func TestTokenTestSuite(t *testing.T) {
	suite.Run(t, new(TokenTestSuite))
}
