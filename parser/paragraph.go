package parser

import "github.com/takaaa220/golang-toy-markdown-parser/ast"

func (p *Parser) paragraph() (ast.Node, error) {
	line := p.next()

	if line == "" {
		return ast.Node{}, ParseError{Message: "invalid paragraph", Line: p.lineCursor, From: 0, To: len(line)}
	}

	return ast.ParagraphNode(line), nil
}
