package parser

import (
	"fmt"
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) unorderedList(currentIndent int, state *blockParsedState) (ast.Node, error) {
	listItems := []ast.Node{}
	var usingSymbol rune
	var beforeListItem *ast.ListItem

	for {
		if !p.hasNext() {
			break
		}

		line := p.peek()
		indent := line.getIndent()
		if indent < currentIndent {
			break
		}
		if indent > currentIndent {
			if beforeListItem == nil {
				return &ast.NodeBase{}, BlockParseError{Message: "invalid unordered list", State: *state}
			}

			children, err := p.Parse(indent)
			if err != nil {
				return &ast.NodeBase{}, err
			}

			beforeListItem.AppendChildren(children...)

			continue
		}

		listText, symbol, isListItem := getUnorderedListItemText(line.getText(indent), usingSymbol)
		if !isListItem {
			break
		}

		usingSymbol = symbol

		if beforeListItem != nil {
			listItems = append(listItems, beforeListItem)
		}

		listItemChildren, err := inline(listText)
		if err != nil {
			return &ast.NodeBase{}, err
		}

		beforeListItem = ast.NewListItem(listItemChildren...)
		p.next(state)
	}

	if beforeListItem != nil {
		listItems = append(listItems, beforeListItem)
	}

	if len(listItems) == 0 {
		return &ast.NodeBase{}, BlockParseError{Message: "invalid unordered list", State: *state}
	}

	return ast.NewUnorderedList(listItems...), nil
}

func getUnorderedListItemText(line string, usingSymbol rune) (string, rune, bool) {
	if usingSymbol != 0 {
		prefix := fmt.Sprintf("%c ", usingSymbol)
		if strings.HasPrefix(line, prefix) {
			return strings.TrimPrefix(line, prefix), usingSymbol, true
		}

		return "", usingSymbol, false
	}

	availableSymbols := []rune{'-', '*', '+'}
	for _, symbol := range availableSymbols {
		prefix := fmt.Sprintf("%c ", symbol)
		if strings.HasPrefix(line, prefix) {
			return strings.TrimPrefix(line, prefix), symbol, true
		}
	}

	return "", 0, false
}

func isUnorderedList(line string) bool {
	availableSymbols := []rune{'-', '*', '+'}
	for _, symbol := range availableSymbols {
		if strings.HasPrefix(line, fmt.Sprintf("%c ", symbol)) {
			return true
		}
	}

	return false
}
