package parser

import (
	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) block(currentIndent int) (ast.Node, error) {
	line := p.peek()[currentIndent:]

	switch {
	case isHeading(line):
		return p.heading(currentIndent)
	case isBlockquote(line):
		return p.blockquote(currentIndent)
	case isCodeblock(line):
		return p.codeblock(currentIndent)
	case isUnorderedList(line):
		return p.unorderedList(currentIndent)
	case isOrderedList(line):
		return p.orderedList(currentIndent)
	// case line[0] == '|':
	// 	return table()
	default:
		return p.paragraph(currentIndent)
	}
}
