package parser

import (
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) codeblock(currentIndent int) (ast.Node, error) {
	firstLine := p.next()[currentIndent:]
	if !strings.HasPrefix(firstLine, "```") {
		return ast.Node{}, ParseError{Message: "invalid codeblock", Line: p.lineCursor, From: 0, To: len(firstLine)}
	}
	// TODO: support language

	lines := []string{}
	for {
		if !p.hasNext() {
			break
		}

		line := p.next()[currentIndent:]
		if strings.HasPrefix(line, "```") {
			break
		}

		lines = append(lines, line)
	}

	return ast.CodeBlockNode(lines), nil
}
