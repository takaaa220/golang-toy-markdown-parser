package renderer

import (
	"fmt"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

type Renderer struct{}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) Render(nodes []ast.Node) string {
	result := ""

	for _, node := range nodes {
		ast.Walk(node, func(node ast.Node) {
			result += r.render(node)
		})
	}

	return result
}

func (r *Renderer) render(node ast.Node) string {
	switch node := node.(type) {
	case *ast.Text:
		return r.text(*node)
	case *ast.Paragraph:
		return r.paragraph(*node)
	case *ast.Heading:
		return r.heading(*node)
	case *ast.CodeBlock:
		return r.codeBlock(*node)
	case *ast.BlockQuote:
		return r.blockQuote(*node)
	case *ast.UnorderedList:
		return r.unorderedList(*node)
	case *ast.OrderedList:
		return r.orderedList(*node)
	case *ast.ListItem:
		return r.listItem(*node)
	case *ast.Table:
		return r.table(*node)
	case *ast.TableRow:
		return r.tableRow(*node)
	case *ast.TableCell:
		return r.tableCell(*node)
	case *ast.Empty:
		return r.empty(*node)
	case *ast.Strong:
		return r.strong(*node)
	case *ast.Italic:
		return r.italic(*node)
	case *ast.StrikeThrough:
		return r.strikeThrough(*node)
	case *ast.Link:
		return r.link(*node)
	case *ast.Image:
		return r.image(*node)
	case *ast.Code:
		return r.code(*node)
	case *ast.NewLine:
		return r.newLine(*node)
	default:
		panic(fmt.Sprintf("unknown node type: %T", node))
	}
}

func (r *Renderer) paragraph(p ast.Paragraph) string {
	return "<p>" + r.Render(p.Children()) + "</p>"
}

func (r *Renderer) heading(h ast.Heading) string {
	return fmt.Sprintf("<h%d>%s</h%d>", h.Level, r.Render(h.Children()), h.Level)
}

func (r *Renderer) codeBlock(cb ast.CodeBlock) string {
	return fmt.Sprintf("<pre><code class=\"language-%s\">%s</code></pre>", cb.Language, cb.Text)
}

func (r *Renderer) blockQuote(bq ast.BlockQuote) string {
	return "<blockquote>" + r.Render(bq.Children()) + "</blockquote>"
}

func (r *Renderer) unorderedList(ul ast.UnorderedList) string {
	return "<ul>" + r.Render(ul.Children()) + "</ul>"
}

func (r *Renderer) orderedList(ol ast.OrderedList) string {
	return "<ol>" + r.Render(ol.Children()) + "</ol>"
}

func (r *Renderer) listItem(li ast.ListItem) string {
	return "<li>" + r.Render(li.Children()) + "</li>"
}

func (r *Renderer) table(t ast.Table) string {
	return "<table>" + r.Render(t.Children()) + "</table>"
}

func (r *Renderer) tableRow(tr ast.TableRow) string {
	return "<tr>" + r.Render(tr.Children()) + "</tr>"
}

func (r *Renderer) tableCell(tc ast.TableCell) string {
	return "<td>" + r.Render(tc.Children()) + "</td>"
}

func (r *Renderer) empty(e ast.Empty) string {
	return ""
}

func (r *Renderer) text(t ast.Text) string {
	return t.Text
}

func (r *Renderer) strong(s ast.Strong) string {
	return "<strong>" + r.Render(s.Children()) + "</strong>"
}

func (r *Renderer) italic(i ast.Italic) string {
	return "<em>" + r.Render(i.Children()) + "</em>"
}

func (r *Renderer) strikeThrough(st ast.StrikeThrough) string {
	return "<del>" + r.Render(st.Children()) + "</del>"
}

func (r *Renderer) link(l ast.Link) string {
	return fmt.Sprintf("<a href=\"%s\">%s</a>", l.Href, r.Render(l.Children()))
}

func (r *Renderer) image(i ast.Image) string {
	return fmt.Sprintf("<img src=\"%s\" alt=\"%s\">", i.Src, i.Alt)
}

func (r *Renderer) code(c ast.Code) string {
	return fmt.Sprintf("<code>%s</code>", c.Text)
}

func (r *Renderer) newLine(n ast.NewLine) string {
	return "<br>"
}
