package ast

import "strconv"

// AST node type declarations
const (
	BYTE string = "BYTE"
	INT  string = "INT"
	LIST string = "LIST"
	DICT string = "DICT"
)

type BaseNodeInterface interface {
	Type() string
}

type ByteNode struct {
	value []byte
}

func NewByteNode(content string) *ByteNode {
	bytes := []byte(content)
	return &ByteNode{
		value: bytes,
	}
}

func (b *ByteNode) Type() string {
	return BYTE
}

func (b *ByteNode) Value() []byte {
	return b.value
}

type IntNode struct {
	value int
}

func NewIntNode(intString string) *IntNode {
	intValue, err := strconv.Atoi(intString)
	if err != nil {
		panic(err)
	}
	return &IntNode{
		value: intValue,
	}
}

func (i *IntNode) Type() string {
	return INT
}

func (i *IntNode) Value() int {
	return i.value
}

type ListNode struct {
	value []BaseNodeInterface
}

func NewListNode() *ListNode {
	return &ListNode{
		value: []BaseNodeInterface{},
	}
}

func (l *ListNode) Type() string {
	return LIST
}

func (l *ListNode) Value() []BaseNodeInterface {
	return l.value
}

func (l *ListNode) Add(n BaseNodeInterface) []BaseNodeInterface {
	l.value = append(l.value, n)
	return l.value
}

type DictNode struct {
	keys  []string
	value map[string]BaseNodeInterface
}

func NewDictNode() *DictNode {
	return &DictNode{
		keys:  []string{},
		value: map[string]BaseNodeInterface{},
	}
}

func (d *DictNode) Type() string {
	return DICT
}

func (d *DictNode) Value() map[string]BaseNodeInterface {
	return d.value
}

func (d *DictNode) Add(keyNode *ByteNode, valueNode BaseNodeInterface) {
	key := string(keyNode.Value())
	d.keys = append(d.keys, key)
	d.value[key] = valueNode
}

func (d *DictNode) Get(keyNode *ByteNode) BaseNodeInterface {
	key := string(keyNode.Value())
	return d.value[key]
}

func (d *DictNode) GetKeys() []string {
	return d.keys
}
