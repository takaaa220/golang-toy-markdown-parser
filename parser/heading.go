package parser

import (
	"regexp"
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) heading(currentIndent int) (ast.Node, error) {
	line := p.next().getText(currentIndent)

	level := 0
	for line[level] == '#' {
		level++
	}
	if level == 0 || level > 6 {
		return ast.Node{}, ParseError{Message: "invalid heading", Line: p.lineCursor, From: 0, To: len(line)}
	}

	if line[level] != ' ' {
		return ast.Node{}, ParseError{Message: "invalid heading", Line: p.lineCursor, From: 0, To: len(line)}
	}

	children, err := inline(strings.TrimLeft(line[level:], " "))
	if err != nil {
		return ast.Node{}, err
	}

	return ast.HeadingNode(level, children...), nil
}

func isHeading(line string) bool {
	return regexp.MustCompile(`^#{1,6} `).MatchString(line)
}
