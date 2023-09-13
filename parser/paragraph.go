package parser

import "github.com/takaaa220/golang-toy-markdown-parser/ast"

func (p *Parser) paragraph(currentIndent int, state *blockParsedState) (ast.Node, error) {
	line := p.next(state).getText(currentIndent)

	if line == "" {
		return ast.NewEmpty(), nil
	}

	children, err := inline(line)
	if err != nil {
		return &ast.NodeBase{}, err
	}

	return ast.NewParagraph(children...), nil
}
