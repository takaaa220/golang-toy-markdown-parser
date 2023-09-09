package parser

import (
	"fmt"
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (p *Parser) unorderedList(currentIndent int) (ast.Node, error) {
	listItems := []ast.Node{}
	var usingSymbol string
	var beforeListItem *ast.Node

	for {
		if !p.hasNext() {
			break
		}

		line, indent := removeIndent(p.peek())
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

		listText, isListItem := getListItemText(line, func(line string) (bool, string) {
			if usingSymbol != "" {
				prefix := fmt.Sprintf("%s ", usingSymbol)
				if strings.HasPrefix(line, prefix) {
					return true, strings.TrimPrefix(line, prefix)
				}

				return false, ""
			}

			availableSymbols := []string{"-", "*", "+"}
			for _, symbol := range availableSymbols {
				prefix := fmt.Sprintf("%s ", symbol)
				if strings.HasPrefix(line, prefix) {
					usingSymbol = symbol

					return true, strings.TrimPrefix(line, prefix)
				}
			}

			return false, ""
		})
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
		return ast.Node{}, ParseError{Message: "invalid unordered list", Line: p.lineCursor, From: 0, To: 1}
	}

	return ast.UnorderedListNode(listItems...), nil
}

func (p *Parser) orderedList(currentIndent int) (ast.Node, error) {
	listItems := []ast.Node{}
	var beforeListItem *ast.Node

	for {
		if !p.hasNext() {
			break
		}

		line, indent := removeIndent(p.peek())
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

		listText, isListItem := getListItemText(line, func(line string) (bool, string) {
			num := len(listItems) + 1
			if beforeListItem != nil {
				num++
			}

			prefix := fmt.Sprintf("%d. ", num)
			if strings.HasPrefix(line, prefix) {
				return true, strings.TrimPrefix(line, prefix)
			}

			return false, ""
		})
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

func getListItemText(line string, checkSymbol func(string) (bool, string)) (string, bool) {
	if isValid, symbolRemovedLine := checkSymbol(line); isValid {
		return symbolRemovedLine, true
	}

	return "", false
}
