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

type HeadingAttribute struct {
	Level int
}

type TableAttribute struct{}

type LinkAttribute struct {
	Href string
}

type ImageAttribute struct {
	Alt string
	Src string
}

type CodeBlockAttribute struct {
	Language string
}

type Node struct {
	Type      NodeType
	Text      string
	Children  []Node
	Attribute interface{}
	// Raw       string
}

func HeadingNode(level int, children ...Node) Node {
	return Node{
		Type:      Heading,
		Children:  children,
		Attribute: HeadingAttribute{Level: level},
	}
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

func TableNode(attribute TableAttribute) Node {
	return Node{Type: Table, Attribute: attribute}
}

func CodeBlockNode(lines []string, language string) Node {
	return Node{
		Type: CodeBlock,
		Children: []Node{
			TextNode(strings.Join(lines, "\n")),
		},
		Attribute: CodeBlockAttribute{Language: language},
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
	return Node{
		Type:      Image,
		Attribute: ImageAttribute{Alt: alt, Src: src},
	}
}

func LinkNode(href string, children ...Node) Node {
	return Node{
		Type:      Link,
		Children:  children,
		Attribute: LinkAttribute{Href: href},
	}
}

func NewLineNode() Node {
	return Node{Type: NewLine}
}

func EscapeNode() Node {
	return Node{Type: Escape}
}
