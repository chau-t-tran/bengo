package ast

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ASTTestSuite struct {
	suite.Suite
}

func (suite *ASTTestSuite) TestByteNode() {
	stringValue := "hello world"
	byteNode := NewByteNode(stringValue)
	assert.Equal(suite.T(), BYTE, byteNode.Type())
	assert.Equal(suite.T(), []byte(stringValue), byteNode.Value())
}

func (suite *ASTTestSuite) TestIntNode() {
	intValue := 20
	stringValue := strconv.Itoa(intValue)
	intNode := NewIntNode(stringValue)
	assert.Equal(suite.T(), INT, intNode.Type())
	assert.Equal(suite.T(), intValue, intNode.Value())
}

func (suite *ASTTestSuite) TestListNode() {
	listNode := NewListNode()

	byteNode := NewByteNode("hello")
	listNode.Add(byteNode)

	intNode := NewIntNode("20")
	listNode.Add(intNode)

	assert.Equal(suite.T(), LIST, listNode.Type())
	assert.Equal(suite.T(), byteNode, listNode.Value()[0])
	assert.Equal(suite.T(), intNode, listNode.Value()[1])
}

func (suite *ASTTestSuite) TestDictNode() {
	dictNode := NewDictNode()

	key1Node := NewByteNode("foo")
	key2Node := NewByteNode("foobar")
	keyList := []string{"foo", "foobar"}
	value1Node := NewByteNode("bar")
	value2Node := NewIntNode("20")

	dictNode.Add(key1Node, value1Node)
	dictNode.Add(key2Node, value2Node)

	assert.Equal(suite.T(), DICT, dictNode.Type())
	assert.Equal(suite.T(), value1Node, dictNode.Get(key1Node))
	assert.Equal(suite.T(), value2Node, dictNode.Get(key2Node))
	assert.Equal(suite.T(), keyList, dictNode.GetKeys())
}

func TestASTTestSuite(t *testing.T) {
	suite.Run(t, new(ASTTestSuite))
}
