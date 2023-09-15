package renderer

import (
	"fmt"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

type Renderer interface {
	Text(node ast.Text) string
	Paragraph(node ast.Paragraph) string
	Heading(node ast.Heading) string
	CodeBlock(node ast.CodeBlock) string
	BlockQuote(node ast.BlockQuote) string
	UnorderedList(node ast.UnorderedList) string
	OrderedList(node ast.OrderedList) string
	ListItem(node ast.ListItem) string
	Table(node ast.Table) string
	TableRow(node ast.TableRow) string
	TableCell(node ast.TableCell) string
	Empty(node ast.Empty) string
	Strong(node ast.Strong) string
	Italic(node ast.Italic) string
	StrikeThrough(node ast.StrikeThrough) string
	Link(node ast.Link) string
	Image(node ast.Image) string
	Code(node ast.Code) string
	NewLine(node ast.NewLine) string
}

type DefaultRenderer struct{}

func NewDefaultRenderer() Renderer {
	return DefaultRenderer{}
}

func (r DefaultRenderer) Paragraph(p ast.Paragraph) string {
	return "<p>" + Render(p.Children(), r) + "</p>"
}

func (r DefaultRenderer) Heading(h ast.Heading) string {
	return fmt.Sprintf("<h%d>%s</h%d>", h.Level, Render(h.Children(), r), h.Level)
}

func (r DefaultRenderer) CodeBlock(cb ast.CodeBlock) string {
	return fmt.Sprintf("<pre><code class=\"language-%s\">%s</code></pre>", cb.Language, cb.Text)
}

func (r DefaultRenderer) BlockQuote(bq ast.BlockQuote) string {
	return "<blockquote>" + Render(bq.Children(), r) + "</blockquote>"
}

func (r DefaultRenderer) UnorderedList(ul ast.UnorderedList) string {
	return "<ul>" + Render(ul.Children(), r) + "</ul>"
}

func (r DefaultRenderer) OrderedList(ol ast.OrderedList) string {
	return "<ol>" + Render(ol.Children(), r) + "</ol>"
}

func (r DefaultRenderer) ListItem(li ast.ListItem) string {
	return "<li>" + Render(li.Children(), r) + "</li>"
}

func (r DefaultRenderer) Table(t ast.Table) string {
	return "<table>" + Render(t.Children(), r) + "</table>"
}

func (r DefaultRenderer) TableRow(tr ast.TableRow) string {
	return "<tr>" + Render(tr.Children(), r) + "</tr>"
}

func (r DefaultRenderer) TableCell(tc ast.TableCell) string {
	return "<td>" + Render(tc.Children(), r) + "</td>"
}

func (r DefaultRenderer) Empty(e ast.Empty) string {
	return ""
}

func (r DefaultRenderer) Text(t ast.Text) string {
	return t.Text
}

func (r DefaultRenderer) Strong(s ast.Strong) string {
	return "<strong>" + Render(s.Children(), r) + "</strong>"
}

func (r DefaultRenderer) Italic(i ast.Italic) string {
	return "<em>" + Render(i.Children(), r) + "</em>"
}

func (r DefaultRenderer) StrikeThrough(st ast.StrikeThrough) string {
	return "<del>" + Render(st.Children(), r) + "</del>"
}

func (r DefaultRenderer) Link(l ast.Link) string {
	return fmt.Sprintf("<a href=\"%s\">%s</a>", l.Href, Render(l.Children(), r))
}

func (r DefaultRenderer) Image(i ast.Image) string {
	return fmt.Sprintf("<img src=\"%s\" alt=\"%s\">", i.Src, i.Alt)
}

func (r DefaultRenderer) Code(c ast.Code) string {
	return fmt.Sprintf("<code>%s</code>", c.Text)
}

func (r DefaultRenderer) NewLine(n ast.NewLine) string {
	return "<br>"
}

func Render(nodes []ast.Node, renderer Renderer) string {
	result := ""

	for _, node := range nodes {
		result += render(node, renderer)
	}

	return result
}

func render(node ast.Node, r Renderer) string {
	switch node := node.(type) {
	case *ast.Text:
		return r.Text(*node)
	case *ast.Paragraph:
		return r.Paragraph(*node)
	case *ast.Heading:
		return r.Heading(*node)
	case *ast.CodeBlock:
		return r.CodeBlock(*node)
	case *ast.BlockQuote:
		return r.BlockQuote(*node)
	case *ast.UnorderedList:
		return r.UnorderedList(*node)
	case *ast.OrderedList:
		return r.OrderedList(*node)
	case *ast.ListItem:
		return r.ListItem(*node)
	case *ast.Table:
		return r.Table(*node)
	case *ast.TableRow:
		return r.TableRow(*node)
	case *ast.TableCell:
		return r.TableCell(*node)
	case *ast.Empty:
		return r.Empty(*node)
	case *ast.Strong:
		return r.Strong(*node)
	case *ast.Italic:
		return r.Italic(*node)
	case *ast.StrikeThrough:
		return r.StrikeThrough(*node)
	case *ast.Link:
		return r.Link(*node)
	case *ast.Image:
		return r.Image(*node)
	case *ast.Code:
		return r.Code(*node)
	case *ast.NewLine:
		return r.NewLine(*node)
	default:
		panic(fmt.Sprintf("unknown node type: %T", node))
	}
}
