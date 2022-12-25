package marshaller

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UnmarshalTestSuite struct {
	suite.Suite
}

func (suite *UnmarshalTestSuite) TestUnmarshalInt() {
	type IntStruct struct {
		Value int `bencode:"value"`
	}

	raw := "d5:valuei42ee"
	var actual IntStruct

	err := Unmarshal(raw, &actual)
	assert.NoError(suite.T(), err)

	expected := IntStruct{Value: 42}

	assert.Equal(suite.T(), expected, actual)
}

func TestUnmarshalTestSuite(t *testing.T) {
	suite.Run(t, new(UnmarshalTestSuite))
}
