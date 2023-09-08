package parser

import (
	"fmt"
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

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
	lines      []string
	lineCursor int
}

func NewParser(input string) *Parser {
	return &Parser{lines: strings.Split(input, "\n"), lineCursor: -1}
}

func (l *Parser) Lex() ([]ast.Node, error) {
	nodes := []ast.Node{}

	for l.hasNext() {
		node, err := l.block(0)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, node)
		l.lineCursor++
	}

	return nodes, nil
}

func (p *Parser) hasNext() bool {
	return p.lineCursor < len(p.lines)-1
}

func (p *Parser) peek() string {
	line := p.lines[p.lineCursor+1]

	return line
}

func (l *Parser) next() {
	l.lineCursor += 1
}

func (p *Parser) back() {
	p.lineCursor -= 1
}
