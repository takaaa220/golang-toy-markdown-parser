package ast

import (
	"strings"
)

type NodeType string

// Block Tokens
const (
	Heading       NodeType = "Heading"
	Paragraph     NodeType = "Paragraph"
	OrderedList   NodeType = "OrderedList"
	UnorderedList NodeType = "UnorderedList"
	ListItem      NodeType = "ListItem"
	Table         NodeType = "Table"
	CodeBlock     NodeType = "CodeBlock"
	BlockQuote    NodeType = "BlockQuote"
	Header        NodeType = "Header"
	Empty         NodeType = "Empty"
)

// Inline Tokens
const (
	Text    NodeType = "Text"
	Strong  NodeType = "Strong"
	Italic  NodeType = "Italic"
	Image   NodeType = "Image"
	Link    NodeType = "Link"
	NewLine NodeType = "NewLine"
	Escape  NodeType = "Escape"
)

type Node struct {
	Type     NodeType
	Text     string
	Level    int
	Raw      string
	Children []Node
}

// for debug
// func inner(n Node, indent int) string {
// 	indentStr := strings.Repeat(" ", indent)
// 	if len(n.Children) == 0 {
// 		return indentStr + string(n.Type) + " " + n.Text
// 	}

// 	children := []string{}
// 	for _, child := range n.Children {
// 		children = append(children, inner(child, indent+2))
// 	}
// 	return indentStr + string(n.Type) + " " + n.Text + "\n" + strings.Join(children, "\n")
// }

// func (n Node) String() string {
// 	return inner(n, 0)
// }

func HeadingNode(level int, text string) Node {
	return Node{Type: Heading, Level: level, Children: []Node{TextNode(text)}}
}

func ParagraphNode(text string) Node {
	return Node{Type: Paragraph, Children: []Node{TextNode(text)}}
}

func UnorderedListNode(listItems ...Node) Node {
	return Node{Type: UnorderedList, Children: listItems}
}

func OrderedListNode(listItems ...Node) Node {
	return Node{Type: OrderedList, Children: listItems}
}

func ListItemNode(children ...Node) Node {
	return Node{
		Type:     ListItem,
		Children: children,
	}
}

func TableNode() Node {
	return Node{Type: Table}
}

func CodeBlockNode(lines []string) Node {
	return Node{
		Type: CodeBlock,
		Children: []Node{
			TextNode(strings.Join(lines, "\n")),
		},
	}
}

func BlockQuoteNode(lines []string) Node {
	return Node{
		Type:     BlockQuote,
		Children: []Node{TextNode(strings.Join(lines, "\n"))},
	}
}

func HeaderNode() Node {
	return Node{Type: Header}
}

func EmptyNode() Node {
	return Node{Type: Empty}
}

func TextNode(text string) Node {
	return Node{Type: Text, Text: text}
}

func StrongNode() Node {
	return Node{Type: Strong}
}

func ItalicNode() Node {
	return Node{Type: Italic}
}

func ImageNode() Node {
	return Node{Type: Image}
}

func LinkNode() Node {
	return Node{Type: Link}
}

func NewLineNode() Node {
	return Node{Type: NewLine}
}

func EscapeNode() Node {
	return Node{Type: Escape}
}
