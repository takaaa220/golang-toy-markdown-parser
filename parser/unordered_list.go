package parser

import (
	"fmt"
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) unorderedList(currentIndent int) (ast.Node, error) {
	state := p.newState()

	listItems := []ast.Node{}
	var usingSymbol rune
	var beforeListItem *ast.Node

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
				return ast.Node{}, BlockParseError{Message: "invalid unordered list", State: *state}
			}

			children, err := p.Parse(indent)
			if err != nil {
				return ast.Node{}, err
			}

			beforeListItem.Children = append(beforeListItem.Children, children...)

			continue
		}

		listText, symbol, isListItem := getUnorderedListItemText(line.getText(indent), usingSymbol)
		if !isListItem {
			break
		}

		usingSymbol = symbol

		if beforeListItem != nil {
			listItems = append(listItems, *beforeListItem)
		}

		listItemChildren, err := inline(listText)
		if err != nil {
			return ast.Node{}, err
		}

		listItem := ast.ListItemNode(listItemChildren...)
		beforeListItem = &listItem

		p.next(state)
	}

	if beforeListItem != nil {
		listItems = append(listItems, *beforeListItem)
	}

	if len(listItems) == 0 {
		return ast.Node{}, BlockParseError{Message: "invalid unordered list", State: *state}
	}

	return ast.UnorderedListNode(listItems...), nil
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
	return line[0] == '-' || line[0] == '*' || line[0] == '+'
}
