package parser

import (
	"fmt"
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) unorderedList(currentIndent int) (ast.Node, error) {
	listItems := []ast.Node{}
	var usingSymbol rune
	var beforeListItem *ast.Node

	for {
		if !p.hasNext() {
			break
		}

		line := p.peek()
		indent := getIndent(line)
		line = line[indent:]
		if indent < currentIndent {
			break
		}
		if indent > currentIndent {
			if beforeListItem == nil {
				return ast.Node{}, ParseError{Message: "invalid unordered list", Line: p.lineCursor, From: 0, To: 1}
			}

			children, err := p.Parse(indent)
			if err != nil {
				return ast.Node{}, err
			}

			beforeListItem.Children = append(beforeListItem.Children, children...)

			continue
		}

		listText, symbol, isListItem := getUnorderedListItemText(line, usingSymbol)
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

		p.next()
	}

	if beforeListItem != nil {
		listItems = append(listItems, *beforeListItem)
	}

	if len(listItems) == 0 {
		return ast.Node{}, ParseError{Message: "invalid unordered list", Line: p.lineCursor, From: 0, To: 1}
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
