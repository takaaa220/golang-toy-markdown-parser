package parser

import "github.com/takaaa220/golang-toy-markdown-parser/ast"

func (l *Parser) heading() (ast.Node, error) {
	line := l.lines[l.lineCursor]
	level := 0

	for line[level] == '#' {
		level++
	}
	if level == 0 {
		return ast.Node{}, ParseError{Message: "invalid heading", Line: l.lineCursor, From: 0, To: len(line)}
	}

	return ast.HeadingNode(level, line[level+1:]), nil
}
