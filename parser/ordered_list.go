package parser

import (
	"fmt"
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) orderedList(currentIndent int, state *blockParsedState) (ast.Node, error) {
	listItems := []ast.Node{}
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
				return &ast.NodeBase{}, BlockParseError{Message: "invalid ordered list", State: *state}
			}

			children, err := p.Parse(indent)
			if err != nil {
				return &ast.NodeBase{}, err
			}

			beforeListItem.AppendChildren(children...)

			continue
		}

		currentListItemNumber := len(listItems)
		if beforeListItem != nil {
			currentListItemNumber++
		}

		listText, isListItem := getOrderedListItemText(line.getText(currentIndent), currentListItemNumber)
		if !isListItem {
			break
		}

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
		return &ast.NodeBase{}, BlockParseError{Message: "invalid ordered list", State: *state}
	}

	return ast.NewOrderedList(listItems...), nil
}

func isOrderedList(line string) bool {
	return strings.HasPrefix(line, "1.")
}

func getOrderedListItemText(line string, currentListItemNumber int) (string, bool) {
	prefix := fmt.Sprintf("%d. ", currentListItemNumber+1)
	if strings.HasPrefix(line, prefix) {
		return strings.TrimPrefix(line, prefix), true
	}

	return "", false
}
