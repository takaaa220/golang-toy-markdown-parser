package parser

import (
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (l *Parser) codeblock(currentIndent int) (ast.Node, error) {
	firstLine := l.next()[currentIndent:]
	if !strings.HasPrefix(firstLine, "```") {
		return ast.Node{}, ParseError{Message: "invalid codeblock", Line: l.lineCursor, From: 0, To: len(firstLine)}
	}
	// TODO: support language

	lines := []string{}
	for {
		if !l.hasNext() {
			break
		}

		line := l.next()[currentIndent:]
		if strings.HasPrefix(line, "```") {
			break
		}

		lines = append(lines, line)
	}

	return ast.CodeBlockNode(lines), nil
}
