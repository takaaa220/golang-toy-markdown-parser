package parser

import (
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) blockquote(currentIndent int) (ast.Node, error) {
	children := []ast.Node{}

	state := p.newState()

	for {
		if !p.hasNext() {
			break
		}
		line := p.peek().getText(currentIndent)
		if !isBlockQuote(line) {
			break
		}

		inlineChildren, err := inline(strings.TrimLeft(line[1:], " "))
		if err != nil {
			return ast.Node{}, err
		}

		children = append(children, inlineChildren...)
		p.next(state)
	}

	if len(children) == 0 {
		return ast.Node{}, BlockParseError{Message: "invalid blockquote", State: *state}
	}

	return ast.BlockQuoteNode(children...), nil
}

func isBlockQuote(line string) bool {
	return strings.HasPrefix(line, "> ")
}
