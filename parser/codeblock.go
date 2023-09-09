package parser

import (
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) codeblock(currentIndent int) (ast.Node, error) {
	firstLine := p.next().getText(currentIndent)
	if !strings.HasPrefix(firstLine, "```") {
		return ast.Node{}, ParseError{Message: "invalid codeblock", Line: p.lineCursor, From: 0, To: len(firstLine)}
	}

	language := strings.Trim(firstLine[3:], " ")

	lines := []string{}
	for {
		if !p.hasNext() {
			break
		}

		line := p.next().getText(currentIndent)
		if isCodeblock(line) {
			break
		}

		lines = append(lines, line)
	}

	return ast.CodeBlockNode(lines, language), nil
}

func isCodeblock(line string) bool {
	return strings.HasPrefix(line, "```")
}
