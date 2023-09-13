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

type BlockParseError struct {
	Message string
	State   blockParsedState
}

func (e BlockParseError) Error() string {
	return fmt.Sprintf("[Block]%s at line %d~%d", e.Message, e.State.from, e.State.from+len(e.State.lines)-1)
}

type InlineParseError struct {
	Message string
	Text    string
	From    int
	To      int
}

func (e InlineParseError) Error() string {
	return fmt.Sprintf(`[Inline] %s in '%s' at %d~%d`, e.Message, e.Text, e.From, e.To)
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

		state := p.newState()
		node, err := p.block(indent, state)
		if err != nil {
			return nil, err
		}

		if node.Type() == ast.EmptyType {
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

func (p *Parser) next(state *blockParsedState) Line {
	if !p.hasNext() {
		panic("no next line")
	}

	p.lineCursor++

	line := p.lines[p.lineCursor]
	state.add(line)

	return line
}

func (p Parser) newState() *blockParsedState {
	return &blockParsedState{
		lines: []Line{},
		from:  p.lineCursor + 1,
	}
}

type blockParsedState struct {
	lines []Line
	from  int
}

func (s *blockParsedState) add(line Line) {
	s.lines = append(s.lines, line)
}
