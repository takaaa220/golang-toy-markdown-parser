package parser

import (
	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) block(currentIndent int, state *blockParsedState) (ast.Node, error) {
	line := p.peek().getText(currentIndent)

	switch {
	case isHeading(line):
		return p.heading(currentIndent, state)
	case isBlockQuote(line):
		return p.blockquote(currentIndent, state)
	case isCodeblock(line):
		return p.codeblock(currentIndent, state)
	case isUnorderedList(line):
		return p.unorderedList(currentIndent, state)
	case isOrderedList(line):
		return p.orderedList(currentIndent, state)
	case isTable(line):
		return p.table(currentIndent, state)
	default:
		return p.paragraph(currentIndent, state)
	}
}
