package parser

import (
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (l *Parser) codeblock() (ast.Node, error) {
	firstLine := l.lines[l.lineCursor]
	if !strings.HasPrefix(firstLine, "```") {
		return ast.Node{}, ParseError{Message: "invalid codeblock", Line: l.lineCursor, From: 0, To: len(firstLine)}
	}
	// TODO: support language

	l.lineCursor++

	lines := []string{}
	for {
		line := l.lines[l.lineCursor]
		if strings.HasPrefix(line, "```") {
			break
		}

		lines = append(lines, line)
		l.lineCursor++
	}

	return ast.CodeBlockNode(lines), nil
}
