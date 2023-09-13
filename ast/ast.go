package ast

import (
	"fmt"
	"strings"
)

type NodeType string

// Block Tokens
const (
	HeadingType       NodeType = "Heading"
	ParagraphType     NodeType = "Paragraph"
	OrderedListType   NodeType = "OrderedList"
	UnorderedListType NodeType = "UnorderedList"
	ListItemType      NodeType = "ListItem"
	TableType         NodeType = "Table"
	TableRowType      NodeType = "TableRow"
	TableCellType     NodeType = "TableCell"
	CodeBlockType     NodeType = "CodeBlock"
	BlockQuoteType    NodeType = "BlockQuote"
	HeaderType        NodeType = "Header"
	EmptyType         NodeType = "Empty"
)

// Inline Tokens
const (
	TextType          NodeType = "Text"
	StrongType        NodeType = "Strong"
	ItalicType        NodeType = "Italic"
	StrikeThroughType NodeType = "StrikeThrough"
	CodeType          NodeType = "Code"
	ImageType         NodeType = "Image"
	LinkType          NodeType = "Link"
	NewLineType       NodeType = "NewLine"
	EscapeType        NodeType = "Escape"
)

type HeadingAttribute struct {
	Level int
}

type Node interface {
	Type() NodeType
	Text() string
	Children() []Node
	AppendChildren(children ...Node)
	Dump() string
}

type NodeBase struct {
	text     string
	nodeType NodeType
	children []Node
}

func NewNodeBase(nodeType NodeType, children ...Node) *NodeBase {
	return &NodeBase{
		nodeType: nodeType,
		children: children,
	}
}

func (n *NodeBase) Type() NodeType {
	return n.nodeType
}

func (n *NodeBase) Children() []Node {
	return n.children
}

func (n *NodeBase) Text() string {
	return n.text
}

func (n NodeBase) Dump() string {
	return ""
}

func (n *NodeBase) AppendChildren(children ...Node) {
	n.children = append(n.children, children...)
}

func (n *NodeBase) String() string {
	return _innerNodeString(n, 0)
}

func _innerNodeString(n Node, indent int) string {
	text := fmt.Sprintf("%s%s %s %s", strings.Repeat(" ", indent), n.Type(), n.Text(), n.Dump())

	for _, child := range n.Children() {
		text += "\n" + _innerNodeString(child, indent+2)
	}

	return text
}
