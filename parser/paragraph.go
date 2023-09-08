package parser

import "github.com/takaaa220/golang-toy-markdown-parser/ast"

func (p *Parser) paragraph(currentIndent int) (ast.Node, error) {
	line := p.next()[currentIndent:]

	if line == "" {
		return ast.EmptyNode(), nil
	}

	return ast.ParagraphNode(line), nil
}
