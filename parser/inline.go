package parser

import (
	"regexp"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func inline(text string) ([]ast.Node, error) {
	cursor := 0
	currentText := ""
	nodes := []ast.Node{}

	for cursor < len(text) {
		switch text[cursor] {
		case '*':
			strongText, cursorOffset, hasStrong := strong(text[cursor:])
			if hasStrong {
				strongNodes, err := inline(strongText)
				if err != nil {
					return nil, err
				}

				if currentText != "" {
					nodes = append(nodes, ast.NewText(currentText))
				}
				nodes = append(nodes, ast.NewStrong(strongNodes...))

				cursor += cursorOffset - 1 // -1 for front "*"
				currentText = ""

				break
			}

			italicText, cursorOffset, hasItalic := italic(text[cursor:])
			if hasItalic {
				italicNodes, err := inline(italicText)
				if err != nil {
					return nil, err
				}

				if currentText != "" {
					nodes = append(nodes, ast.NewText(currentText))
				}
				nodes = append(nodes, ast.NewItalic(italicNodes...))

				cursor += cursorOffset - 1 // -1 for front "*"
				currentText = ""

				break
			}

			currentText += string(text[cursor])
		case '~':
			// strike through
			strikeThroughText, cursorOffset, hasStrikeThrough := strikeThrough(text[cursor:])
			if !hasStrikeThrough {
				currentText += string(text[cursor])
				break
			}

			strikeThroughNodes, err := inline(strikeThroughText)
			if err != nil {
				return nil, err
			}

			if currentText != "" {
				nodes = append(nodes, ast.NewText(currentText))
			}
			nodes = append(nodes, ast.NewStrikeThrough(strikeThroughNodes...))

			cursor += cursorOffset - 1 // -1 for front "~"
			currentText = ""
		case '`':
			// code
			codeText, cursorOffset, hasCode := code(text[cursor:])
			if !hasCode {
				currentText += string(text[cursor])
				break
			}

			if currentText != "" {
				nodes = append(nodes, ast.NewText(currentText))
			}
			nodes = append(nodes, ast.NewCode(codeText))

			cursor += cursorOffset - 1 // -1 for front backquote
			currentText = ""
		case '[':
			// link [Example](https://example.com)
			linkText, href, cursorOffset, hasLink := link(text[cursor:])
			if !hasLink {
				currentText += string(text[cursor])
				break
			}

			if currentText != "" {
				nodes = append(nodes, ast.NewText(currentText))
			}
			nodes = append(nodes, ast.NewLink(href, ast.NewText(linkText)))

			cursor += cursorOffset - 1 // -1 for front `[`
			currentText = ""
		case '!':
			// image ![Example](https://example.com)
			alt, src, cursorOffset, hasImage := image(text[cursor:])
			if !hasImage {
				currentText += string(text[cursor])
				break
			}

			if currentText != "" {
				nodes = append(nodes, ast.NewText(currentText))
			}
			nodes = append(nodes, ast.NewImage(alt, src))

			cursor += cursorOffset - 1 // -1 for front `!`
			currentText = ""
		default:
			currentText += string(text[cursor])
		}

		cursor++
	}

	if currentText != "" {
		nodes = append(nodes, ast.NewText(currentText))
	}

	if len(nodes) == 0 {
		return []ast.Node{ast.NewEmpty()}, nil
	}

	return nodes, nil
}

func strong(text string) (string, int, bool) {
	regexp := regexp.MustCompile(`^\*\*(.+)\*\*(.*)$`)
	matches := regexp.FindStringSubmatch(text)
	if len(matches) != 3 {
		return "", 0, false
	}

	return matches[1], len(matches[1]) + 4, true
}

func italic(text string) (string, int, bool) {
	regexp := regexp.MustCompile(`^\*(.+)\*(.*)$`)
	matches := regexp.FindStringSubmatch(text)
	if len(matches) != 3 {
		return "", 0, false
	}

	return matches[1], len(matches[1]) + 2, true
}

func code(text string) (string, int, bool) {
	regexp := regexp.MustCompile("^`(.+)`(.*)$")
	matches := regexp.FindStringSubmatch(text)
	if len(matches) != 3 {
		return "", 0, false
	}

	return matches[1], len(matches[1]) + 2, true
}

func strikeThrough(text string) (string, int, bool) {
	regex := regexp.MustCompile("^~~(.+)~~(.*)$")
	matches := regex.FindStringSubmatch(text)
	if len(matches) == 3 {
		return matches[1], len(matches[1]) + 4, true
	}

	regex = regexp.MustCompile("^~(.+)~(.*)$")
	matches = regex.FindStringSubmatch(text)
	if len(matches) == 3 {
		return matches[1], len(matches[1]) + 2, true
	}

	return "", 0, false
}

func link(text string) (linkText string, href string, cursorOffset int, hasLink bool) {
	regexp := regexp.MustCompile(`^\[(.+)\]\((.+)\)(.*)$`)
	matches := regexp.FindStringSubmatch(text)
	if len(matches) != 4 {
		return "", "", 0, false
	}

	return matches[1], matches[2], len(matches[1]) + len(matches[2]) + 4, true
}

func image(text string) (alt string, src string, cursorOffset int, hasLink bool) {
	regexp := regexp.MustCompile(`^!\[(.*)\]\((.*)\)(.*)$`)
	matches := regexp.FindStringSubmatch(text)
	if len(matches) != 4 {
		return "", "", 0, false
	}

	return matches[1], matches[2], len(matches[1]) + len(matches[2]) + 5, true
}
