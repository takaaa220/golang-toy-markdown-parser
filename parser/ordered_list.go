package parser

import (
	"fmt"
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) orderedList(currentIndent int) (ast.Node, error) {
	listItems := []ast.Node{}
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
				return ast.Node{}, ParseError{Message: "invalid ordered list", Line: p.lineCursor, From: 0, To: 1}
			}

			children, err := p.Parse(indent)
			if err != nil {
				return ast.Node{}, err
			}

			beforeListItem.Children = append(beforeListItem.Children, children...)

			continue
		}

		currentListItemNumber := len(listItems)
		if beforeListItem != nil {
			currentListItemNumber++
		}

		listText, isListItem := getOrderedListItemText(line, currentListItemNumber)
		if !isListItem {
			break
		}

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
		return ast.Node{}, ParseError{Message: "invalid ordered list", Line: p.lineCursor, From: 0, To: 1}
	}

	return ast.OrderedListNode(listItems...), nil
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
