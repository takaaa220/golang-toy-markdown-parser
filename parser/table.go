package parser

import (
	"regexp"
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) table(currentIndent int, state *blockParsedState) (ast.Node, error) {
	if !p.hasNext() {
		return ast.Node{}, BlockParseError{Message: "invalid table", State: *state}
	}

	headerCellTexts := getCellTexts(p.next(state).getText(currentIndent))
	columnLength := len(headerCellTexts)

	headerRow, err := convertToRow(headerCellTexts, columnLength)
	if err != nil {
		return ast.Node{}, err
	}

	if !p.hasNext() {
		return ast.Node{}, BlockParseError{Message: "invalid table", State: *state}
	}

	aligns, ok := convertAligns(getCellTexts(p.next(state).getText(currentIndent)))
	if !ok {
		return ast.Node{}, BlockParseError{Message: "invalid table", State: *state}
	}

	if len(aligns) != columnLength {
		return ast.Node{}, BlockParseError{Message: "invalid table", State: *state}
	}

	columnDefinitions := make([]ast.TableColumnDefinition, columnLength)
	for i := 0; i < columnLength; i++ {
		columnDefinitions[i] = ast.TableColumnDefinition{
			Align: aligns[i],
		}
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

		p.next(state)

		cellTexts := getCellTexts(line)
		if len(cellTexts) > columnLength {
			return ast.Node{}, BlockParseError{Message: "invalid table", State: *state}
		}

		row, err := convertToRow(cellTexts, columnLength)
		if err != nil {
			return ast.Node{}, err
		}

		rows = append(rows, row)
	}

	if len(rows) == 0 {
		return ast.Node{}, BlockParseError{Message: "invalid table", State: *state}
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

func convertToRow(cellTexts []string, columnLength int) (ast.Node, error) {
	cells := make([]ast.Node, columnLength)

	for i := 0; i < columnLength; i++ {
		cellText := ""
		if i < len(cellTexts) {
			cellText = cellTexts[i]
		}

		inlineChildren, err := inline(cellText)
		if err != nil {
			return ast.Node{}, err
		}

		cells[i] = ast.TableCellNode(inlineChildren...)
	}

	return ast.TableRowNode(cells...), nil
}

func isTable(line string) bool {
	return line[0] == '|'
}
