package parser

import (
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (l *Parser) blockquote() (ast.Node, error) {
	lines := []string{}

	for {
		line := l.lines[l.lineCursor]
		if line[0] != '>' {
			l.lineCursor--
			break
		}

		lines = append(lines, strings.TrimLeft(line[1:], " "))
		l.lineCursor++
	}

	if len(lines) == 0 {
		return ast.Node{}, ParseError{Message: "invalid blockquote", Line: l.lineCursor, From: 0, To: len(lines[l.lineCursor])}
	}

	return ast.BlockQuoteNode(lines), nil
}
