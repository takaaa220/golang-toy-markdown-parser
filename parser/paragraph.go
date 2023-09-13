package parser

import "github.com/takaaa220/golang-toy-markdown-parser/ast"

func (p *Parser) paragraph(currentIndent int, state *blockParsedState) (ast.Node, error) {
	line := p.next(state).getText(currentIndent)

	if line == "" {
		return ast.EmptyNode(), nil
	}

	children, err := inline(line)
	if err != nil {
		return ast.Node{}, err
	}

	return ast.ParagraphNode(children...), nil
}
