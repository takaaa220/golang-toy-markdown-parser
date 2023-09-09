package parser

import (
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) heading(currentIndent int) (ast.Node, error) {
	line := p.next()[currentIndent:]
	level := 0

	for line[level] == '#' {
		level++
	}
	if level == 0 {
		return ast.Node{}, ParseError{Message: "invalid heading", Line: p.lineCursor, From: 0, To: len(line)}
	}

	headingText := line[level:]
	if !strings.HasPrefix(headingText, " ") {
		return ast.Node{}, ParseError{Message: "invalid heading", Line: p.lineCursor, From: 0, To: len(line)}
	}

	children, err := inline(strings.TrimLeft(headingText, " "))
	if err != nil {
		return ast.Node{}, err
	}

	return ast.HeadingNode(level, children...), nil
}
