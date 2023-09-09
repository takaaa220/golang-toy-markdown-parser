package parser

import (
	"regexp"
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) block(currentIndent int) (ast.Node, error) {
	line := p.peek()[currentIndent:]

	switch {
	case line[0] == '#':
		return p.heading(currentIndent)
	case line[0] == '>':
		return p.blockquote(currentIndent)
	case strings.HasPrefix(line, "```"):
		return p.codeblock(currentIndent)
	case line[0] == '-' || line[0] == '+' || line[0] == '*':
		return p.unorderedList(currentIndent)
	case regexp.MustCompile(`^\d+\.`).MatchString(line):
		return p.orderedList(currentIndent)
	// case line[0] == '|':
	// 	return table()
	default:
		return p.paragraph(currentIndent)
	}
}
