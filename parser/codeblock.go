package parser

import (
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) codeblock(currentIndent int) (ast.Node, error) {
	state := p.newState()

	firstLine := p.next(state).getText(currentIndent)
	if !strings.HasPrefix(firstLine, "```") {
		return ast.Node{}, BlockParseError{Message: "invalid codeblock", State: *state}
	}

	language := strings.Trim(firstLine[3:], " ")

	lines := []string{}
	for {
		if !p.hasNext() {
			break
		}

		line := p.next(state).getText(currentIndent)
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
