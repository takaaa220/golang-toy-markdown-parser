package parser

import (
	"regexp"
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) heading(currentIndent int, state *blockParsedState) (ast.Node, error) {
	line := p.next(state).getText(currentIndent)

	level := 0
	for line[level] == '#' {
		level++
	}
	if level == 0 || level > 6 {
		return &ast.NodeBase{}, BlockParseError{Message: "invalid heading", State: *state}
	}

	if line[level] != ' ' {
		return &ast.NodeBase{}, BlockParseError{Message: "invalid heading", State: *state}
	}

	children, err := inline(strings.TrimLeft(line[level:], " "))
	if err != nil {
		return &ast.NodeBase{}, err
	}

	return ast.NewHeading(level, children...), nil
}

func isHeading(line string) bool {
	return regexp.MustCompile(`^#{1,6} `).MatchString(line)
}
