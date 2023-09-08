package parser

import (
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (l *Parser) heading(currentIndent int) (ast.Node, error) {
	line := l.next()[currentIndent:]
	level := 0

	for line[level] == '#' {
		level++
	}
	if level == 0 {
		return ast.Node{}, ParseError{Message: "invalid heading", Line: l.lineCursor, From: 0, To: len(line)}
	}

	headingText := line[level:]
	if !strings.HasPrefix(headingText, " ") {
		return ast.Node{}, ParseError{Message: "invalid heading", Line: l.lineCursor, From: 0, To: len(line)}
	}

	return ast.HeadingNode(level, strings.TrimLeft(headingText, " ")), nil
}
