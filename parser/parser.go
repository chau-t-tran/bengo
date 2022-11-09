package parser

import (
	"errors"

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
