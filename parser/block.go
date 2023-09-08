package parser

import (
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (l *Parser) block(currentIndent int) (ast.Node, error) {
	line := l.lines[l.lineCursor][currentIndent:]
	if line == "" {
		return ast.EmptyNode(), nil
	}

	switch {
	case line[0] == '#':
		return l.heading()
	case line[0] == '>':
		return l.blockquote()
	case strings.HasPrefix(line, "```"):
		return l.codeblock()
	case line[0] == '-' || line[0] == '+' || line[0] == '*':
		return l.unorderedList(0)
	// case strings.HasPrefix(line, "1."):
	// 	return orderedList()
	// case line[0] == '|':
	// 	return table()
	default:
		return l.paragraph()
	}
}
