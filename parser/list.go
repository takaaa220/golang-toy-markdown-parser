package parser

import (
	"fmt"
	"strings"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func (l *Parser) unorderedList(currentIndent int) (ast.Node, error) {
	listItems := []ast.Node{}
	var usingSymbol string
	var beforeListItem *ast.Node

	for {
		if !l.hasNext() {
			break
		}

		isListItem, indent, listText := decomposeLine(l.peek(), func(line string) (bool, string) {
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
		if !isListItem || indent < currentIndent {
			break
		}

		if indent > currentIndent {
			if beforeListItem == nil {
				return ast.Node{}, ParseError{Message: "invalid unordered list", Line: l.lineCursor, From: 0, To: 1}
			}

			items, err := l.unorderedList(indent)
			if err != nil {
				return ast.Node{}, err
			}

			beforeListItem.Children = append(beforeListItem.Children, items)
			continue
		}

		if beforeListItem != nil {
			listItems = append(listItems, *beforeListItem)
		}

		listItem := ast.ListItemNode(ast.TextNode(listText))
		beforeListItem = &listItem

		l.next()
	}

	if beforeListItem != nil {
		listItems = append(listItems, *beforeListItem)
	}

	if len(listItems) == 0 {
		return ast.Node{}, ParseError{Message: "invalid unordered list", Line: l.lineCursor, From: 0, To: 1}
	}

	return ast.UnorderedListNode(listItems...), nil
}

func (l *Parser) orderedList(currentIndent int) (ast.Node, error) {
	listItems := []ast.Node{}
	var beforeListItem *ast.Node

	for {
		if !l.hasNext() {
			break
		}

		isListItem, indent, listText := decomposeLine(l.peek(), func(line string) (bool, string) {
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
		if !isListItem || indent < currentIndent {
			break
		}

		if indent > currentIndent {
			if beforeListItem == nil {
				return ast.Node{}, ParseError{Message: "invalid ordered list", Line: l.lineCursor, From: 0, To: 1}
			}

			items, err := l.orderedList(indent)
			if err != nil {
				return ast.Node{}, err
			}

			beforeListItem.Children = append(beforeListItem.Children, items)
			continue
		}

		if beforeListItem != nil {
			listItems = append(listItems, *beforeListItem)
		}

		listItem := ast.ListItemNode(ast.TextNode(listText))
		beforeListItem = &listItem

		l.next()
	}

	if beforeListItem != nil {
		listItems = append(listItems, *beforeListItem)
	}

	if len(listItems) == 0 {
		return ast.Node{}, ParseError{Message: "invalid ordered list", Line: l.lineCursor, From: 0, To: 1}
	}

	return ast.OrderedListNode(listItems...), nil
}

func decomposeLine(line string, checkSymbol func(string) (bool, string)) (bool, int, string) {
	indent := getIndent(line)
	line = line[indent:]

	if isValid, symbolRemovedLine := checkSymbol(line); isValid {
		return true, indent, symbolRemovedLine
	}

	return false, 0, ""
}

func getIndent(line string) int {
	indent := 0
	for _, c := range line {
		switch c {
		case ' ', '\t':
			indent++
		default:
			return indent
		}
	}

	return indent
}
