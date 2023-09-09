package parser

import (
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) blockquote(currentIndent int) (ast.Node, error) {
	children := []ast.Node{}

	for {
		if !p.hasNext() {
			break
		}
		line := p.peek()[currentIndent:]
		if line[0] != '>' {
			break
		}

		inlineChildren, err := inline(strings.TrimLeft(line[1:], " "))
		if err != nil {
			return ast.Node{}, err
		}

		children = append(children, inlineChildren...)
		p.next()
	}

	if len(children) == 0 {
		return ast.Node{}, ParseError{Message: "invalid blockquote", Line: p.lineCursor, From: 0, To: 0}
	}

	return ast.BlockQuoteNode(children...), nil
}
