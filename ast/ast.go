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
	Text          NodeType = "Text"
	Strong        NodeType = "Strong"
	Italic        NodeType = "Italic"
	StrikeThrough NodeType = "StrikeThrough"
	Code          NodeType = "Code"
	Image         NodeType = "Image"
	Link          NodeType = "Link"
	NewLine       NodeType = "NewLine"
	Escape        NodeType = "Escape"
)

type Node struct {
	Type     NodeType
	Text     string
	Level    int
	Href     string
	Alt      string
	Src      string
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

func HeadingNode(level int, children ...Node) Node {
	return Node{Type: Heading, Level: level, Children: children}
}

func ParagraphNode(children ...Node) Node {
	return Node{Type: Paragraph, Children: children}
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

func BlockQuoteNode(children ...Node) Node {
	return Node{
		Type:     BlockQuote,
		Children: children,
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

func StrongNode(children ...Node) Node {
	return Node{Type: Strong, Children: children}
}

func ItalicNode(children ...Node) Node {
	return Node{Type: Italic, Children: children}
}

func CodeNode(text string) Node {
	return Node{Type: Code, Text: text}
}

func StrikeThroughNode(children ...Node) Node {
	return Node{Type: StrikeThrough, Children: children}
}

func ImageNode(alt string, src string) Node {
	return Node{Type: Image, Alt: alt, Src: src}
}

func LinkNode(href string, children ...Node) Node {
	return Node{Type: Link, Href: href, Children: children}
}

func NewLineNode() Node {
	return Node{Type: NewLine}
}

func EscapeNode() Node {
	return Node{Type: Escape}
}
