package parser

import (
	"regexp"
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (l *Parser) block(currentIndent int) (ast.Node, error) {
	line := l.peek()[currentIndent:]

	switch {
	case line[0] == '#':
		return l.heading(currentIndent)
	case line[0] == '>':
		return l.blockquote(currentIndent)
	case strings.HasPrefix(line, "```"):
		return l.codeblock(currentIndent)
	case line[0] == '-' || line[0] == '+' || line[0] == '*':
		return l.unorderedList(currentIndent)
	case regexp.MustCompile(`^\d+\.`).MatchString(line):
		return l.orderedList(currentIndent)
	// case line[0] == '|':
	// 	return table()
	default:
		return l.paragraph(currentIndent)
	}
}
