package parser

import (
	"regexp"
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) table(currentIndent int) (ast.Node, error) {
	if !p.hasNext() {
		return ast.Node{}, ParseError{Message: "invalid table", Line: p.lineCursor, From: 0, To: 0}
	}

	headerCellTexts := getCellTexts(p.next().getText(currentIndent))
	columnLength := len(headerCellTexts)

	headerCells := make([]ast.Node, columnLength)
	for i, cellText := range headerCellTexts {
		inlineChildren, err := inline(cellText)
		if err != nil {
			return ast.Node{}, err
		}

		headerCells[i] = ast.TableCellNode(inlineChildren...)
	}

	headerRow := ast.TableRowNode(headerCells...)

	if !p.hasNext() {
		return ast.Node{}, ParseError{Message: "invalid table", Line: p.lineCursor, From: 0, To: 0}
	}

	aligns, ok := convertAligns(getCellTexts(p.next().getText(currentIndent)))
	if !ok {
		return ast.Node{}, ParseError{Message: "invalid table", Line: p.lineCursor, From: 0, To: 0}
	}

	columnDefinitions := make([]ast.TableColumnDefinition, columnLength)
	for i := 0; i < columnLength; i++ {
		columnDefinitions[i] = ast.TableColumnDefinition{
			Align: aligns[i],
		}
	}

	if len(columnDefinitions) != columnLength {
		return ast.Node{}, ParseError{Message: "invalid table", Line: p.lineCursor, From: 0, To: 0}
	}

	rows := []ast.Node{}
	for {
		if !p.hasNext() {
			break
		}

		line := p.peek().getText(currentIndent)
		if !isTable(line) {
			break
		}

		cellTexts := getCellTexts(line)
		cells := make([]ast.Node, columnLength)
		for i, cellText := range cellTexts {
			inlineChildren, err := inline(cellText)
			if err != nil {
				return ast.Node{}, err
			}

			cells[i] = ast.TableCellNode(inlineChildren...)
		}

		if len(cells) != columnLength {
			return ast.Node{}, ParseError{Message: "invalid table", Line: p.lineCursor, From: 0, To: 0}
		}

		rows = append(rows, ast.TableRowNode(cells...))

		p.next()
	}

	if len(rows) == 0 {
		return ast.Node{}, ParseError{Message: "invalid table", Line: p.lineCursor, From: 0, To: 0}
	}

	return ast.TableNode(columnDefinitions, append([]ast.Node{headerRow}, rows...)...), nil
}

func getCellTexts(line string) []string {
	split := strings.Split(line, "|")

	cells := make([]string, len(split)-2)
	for i, cell := range split[1 : len(split)-1] {
		cells[i] = strings.Trim(cell, " ")
	}

	return cells
}

func convertAligns(aligns []string) ([]ast.TableColumnAlign, bool) {
	result := make([]ast.TableColumnAlign, len(aligns))

	for i, align := range aligns {
		regexp := regexp.MustCompile(`^(:?)-+(:?)-(:?)$`)
		if !regexp.MatchString(align) {
			return nil, false
		}

		switch {
		case strings.HasPrefix(align, ":-") && strings.HasSuffix(align, "-:"):
			result[i] = ast.TableColumnAlignCenter
		case strings.HasSuffix(align, "-:"):
			result[i] = ast.TableColumnAlignRight
		default:
			result[i] = ast.TableColumnAlignLeft
		}
	}

	return result, true
}

func isTable(line string) bool {
	return line[0] == '|'
}
