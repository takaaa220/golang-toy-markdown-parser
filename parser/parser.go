package parser

import (
	"fmt"
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

type Line struct {
	text string
}

func (l Line) getIndent() int {
	indent := 0
	for _, c := range l.text {
		switch c {
		case ' ', '\t':
			indent++
		default:
			return indent
		}
	}

	return indent
}

func (l Line) getText(indent int) string {
	if l.getIndent() < indent {
		panic("invalid indent")
	}

	return l.text[indent:]
}

type ParseError struct {
	Message string
	Line    int
	From    int
	To      int
}

func (e ParseError) Error() string {
	return fmt.Sprintf("%s in line %d at %d~%d", e.Message, e.Line, e.From, e.To)
}

type Parser struct {
	lines      []Line
	lineCursor int
}

func NewParser(input string) *Parser {
	lines := []Line{}
	for _, line := range strings.Split(input, "\n") {
		lines = append(lines, Line{text: line})
	}

	return &Parser{lines: lines, lineCursor: -1}
}

func (p *Parser) Parse(currentIndent int) ([]ast.Node, error) {
	nodes := []ast.Node{}

	for p.hasNext() {
		indent := p.peek().getIndent()
		if indent < currentIndent {
			break
		}

		node, err := p.block(indent)
		if err != nil {
			return nil, err
		}

		if node.Type == ast.Empty {
			continue
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}

func (p *Parser) hasNext() bool {
	return p.lineCursor < len(p.lines)-1
}

func (p *Parser) peek() Line {
	if !p.hasNext() {
		panic("no next line")
	}

	return p.lines[p.lineCursor+1]
}

func (p *Parser) next() Line {
	if !p.hasNext() {
		panic("no next line")
	}

	p.lineCursor++
	return p.lines[p.lineCursor]
}
