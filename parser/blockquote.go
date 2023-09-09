package parser

import (
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (l *Parser) blockquote(currentIndent int) (ast.Node, error) {
	children := []ast.Node{}

	for {
		if !l.hasNext() {
			break
		}
		line := l.peek()[currentIndent:]
		if line[0] != '>' {
			break
		}

		inlineChildren, err := inline(strings.TrimLeft(line[1:], " "))
		if err != nil {
			return ast.Node{}, err
		}

		children = append(children, inlineChildren...)
		l.next()
	}

	if len(children) == 0 {
		return ast.Node{}, ParseError{Message: "invalid blockquote", Line: l.lineCursor, From: 0, To: 0}
	}

	return ast.BlockQuoteNode(children...), nil
}
