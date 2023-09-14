package ast

import "strings"

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

type Heading struct {
	Node
	Level int
}

func NewHeading(level int, children ...Node) *Heading {
	return &Heading{
		Node:  NewNodeBase(HeadingType, children...),
		Level: level,
	}
}

type Paragraph struct {
	Node
}

func NewParagraph(children ...Node) *Paragraph {
	return &Paragraph{
		Node: NewNodeBase(ParagraphType, children...),
	}
}

type OrderedList struct {
	Node
}

func NewOrderedList(children ...Node) *OrderedList {
	return &OrderedList{
		Node: NewNodeBase(OrderedListType, children...),
	}
}

type UnorderedList struct {
	Node
}

func NewUnorderedList(children ...Node) *UnorderedList {
	return &UnorderedList{
		Node: NewNodeBase(UnorderedListType, children...),
	}
}

type ListItem struct {
	Node
}

func NewListItem(children ...Node) *ListItem {
	return &ListItem{
		Node: NewNodeBase(ListItemType, children...),
	}
}

type TableColumnAlign string

const (
	TableColumnAlignLeft   TableColumnAlign = "Left"
	TableColumnAlignCenter TableColumnAlign = "Center"
	TableColumnAlignRight  TableColumnAlign = "Right"
)

type TableColumnDefinition struct {
	Align TableColumnAlign
}

type TableAttribute struct {
	Columns []TableColumnDefinition
}

type Table struct {
	Node
	Columns []TableColumnDefinition
}

func NewTable(columnDefinitions []TableColumnDefinition, children ...Node) *Table {
	return &Table{
		Node:    NewNodeBase(TableType, children...),
		Columns: columnDefinitions,
	}
}

type TableRow struct {
	Node
}

func NewTableRow(children ...Node) *TableRow {
	return &TableRow{
		Node: NewNodeBase(TableRowType, children...),
	}
}

type TableCell struct {
	Node
}

func NewTableCell(children ...Node) *TableCell {
	return &TableCell{
		Node: NewNodeBase(TableCellType, children...),
	}
}

type BlockQuote struct {
	Node
}

func NewBlockQuote(children ...Node) *BlockQuote {
	return &BlockQuote{
		Node: NewNodeBase(BlockQuoteType, children...),
	}
}

type CodeBlock struct {
	Node
	Language string
	Text     string
}

func NewCodeBlock(lines []string, language string) *CodeBlock {
	return &CodeBlock{
		Node:     NewNodeBase(CodeBlockType),
		Language: language,
		Text:     strings.Join(lines, "\n"),
	}
}

type Empty struct {
	Node
}

func NewEmpty() *Empty {
	return &Empty{
		Node: NewNodeBase(EmptyType),
	}
}
