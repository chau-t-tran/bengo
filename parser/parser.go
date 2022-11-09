package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/chau-t-tran/bengo/ast"
	"github.com/chau-t-tran/bengo/lexer"
	"github.com/chau-t-tran/bengo/token"
)

// TODO: Add custom syntax errors and token expecting mechanism

type parser struct {
	lexer lexer.Lexer
	index int

	current token.Token
	next    token.Token
}

func newParser(lexer lexer.Lexer) (*parser, error) {
	next, err := lexer.NextToken()
	if err != nil {
		return nil, err
	}

	return &parser{
		lexer:   lexer,
		index:   0,
		current: token.NewToken(token.NULL, ""),
		next:    next,
	}, nil
}

func (p *parser) NextToken() (token.Token, error) {
	if p.current.Type == token.TERMINATE {
		return p.current, errors.New("End of input")
	}

	p.index += len(p.next.Literal)
	p.current = p.next
	next, err := p.lexer.NextToken()
	p.next = next
	return p.current, err
}

func (p *parser) PeekToken() token.Token {
	return p.next
}

func (p *parser) parseByte() (node *ast.ByteNode, err error) {
	// Re: Add syntax errors - Expect colon here
	_, err = p.NextToken()
	if err != nil {
		return node, err
	}

	// Re: Add syntax errors - Expect colon here
	_, err = p.NextToken()
	if err != nil {
		return node, err
	}

	bytes, err := p.NextToken()
	if err != nil {
		return node, err
	}

	node = ast.NewByteNode(bytes.Literal)
	return
}

func (p *parser) parseInt() (node *ast.IntNode, err error) {
	// Re: Add syntax errors - Expect "i" here
	_, err = p.NextToken()
	if err != nil {
		return node, err
	}

	integer, err := p.NextToken()
	if err != nil {
		return node, err
	}

	// Re: Add syntax errors - Expect "e" here
	_, err = p.NextToken()
	if err != nil {
		return node, err
	}

	node = ast.NewIntNode(integer.Literal)
	return
}

func (p *parser) parseList() (node *ast.ListNode, err error) {
	node = ast.NewListNode()

	// Re: Add syntax errors - Expect "l" here
	_, err = p.NextToken()

	for p.next.Literal != "e" {
		value, err := p.parseUnknown()
		if err != nil {
			return node, err
		}
		node.Add(value)
	}

	// Re: Add syntax errors - Expect "e" here
	_, err = p.NextToken()
	return
}

func (p *parser) parseUnknown() (node ast.BaseNodeInterface, err error) {
	switch c := p.next.Literal; c {
	case "i":
		return p.parseInt()
	case "l":
		return p.parseList()
	case "e":
		return node, errors.New(
			fmt.Sprintf("End of value"),
		)
	default:
		_, err := strconv.Atoi(c)
		if err != nil {
			return node, errors.New(
				fmt.Sprintf("Unknown symbol %s", c),
			)
		}
		return p.parseByte()
	}
}
