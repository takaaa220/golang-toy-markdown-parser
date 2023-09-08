package parser

import (
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (l *Parser) blockquote(currentIndent int) (ast.Node, error) {
	lines := []string{}

	for {
		if !l.hasNext() {
			break
		}
		line := l.peek()[currentIndent:]
		if line[0] != '>' {
			break
		}

		lines = append(lines, strings.TrimLeft(line[1:], " "))
		l.next()
	}

	if len(lines) == 0 {
		return ast.Node{}, ParseError{Message: "invalid blockquote", Line: l.lineCursor, From: 0, To: len(lines[l.lineCursor])}
	}

	return ast.BlockQuoteNode(lines), nil
}
