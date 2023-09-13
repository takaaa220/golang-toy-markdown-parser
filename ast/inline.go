package ast

type Text struct {
	Node
}

func NewText(text string) *Text {
	return &Text{
		Node: &NodeBase{
			nodeType: TextType,
			text:     text,
		},
	}
}

type Strong struct {
	Node
}

func NewStrong(children ...Node) *Strong {
	return &Strong{
		Node: NewNodeBase(StrongType, children...),
	}
}

type Italic struct {
	Node
}

func NewItalic(children ...Node) *Italic {
	return &Italic{
		Node: NewNodeBase(ItalicType, children...),
	}
}

type StrikeThrough struct {
	Node
}

func NewStrikeThrough(children ...Node) *StrikeThrough {
	return &StrikeThrough{
		Node: NewNodeBase(StrikeThroughType, children...),
	}
}

type Code struct {
	Node
}

func NewCode(text string) *Code {
	return &Code{
		Node: NewNodeBase(CodeType, NewText(text)),
	}
}

type Image struct {
	Node
	Alt string
	Src string
}

func NewImage(alt, src string) *Image {
	return &Image{
		Node: NewNodeBase(ImageType),
		Alt:  alt,
		Src:  src,
	}
}

type Link struct {
	Node
	Href string
}

func NewLink(href string, children ...Node) *Link {
	return &Link{
		Node: NewNodeBase(LinkType, children...),
		Href: href,
	}
}

type NewLine struct {
	Node
}

func NewNewLine() *NewLine {
	return &NewLine{
		Node: NewNodeBase(NewLineType),
	}
}

type Escape struct {
	Node
}

func NewEscape() *Escape {
	return &Escape{
		Node: NewNodeBase(EscapeType),
	}
}
