package renderer

import (
	"strings"
	"testing"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func TestRender(t *testing.T) {
	tests := []struct {
		name  string
		nodes []ast.Node
		want  string
	}{
		{
			name: "test2",
			nodes: []ast.Node{
				ast.NewHeading(1, ast.NewText("heading")),
				ast.NewParagraph(
					ast.NewCode("paragraph"),
					ast.NewText(": "),
					ast.NewLink(
						"https://example.com",
						ast.NewText("link"),
					),
					ast.NewText("!"),
					ast.NewNewLine(),
					ast.NewImage("awesome image", "https://example.com/image.png"),
				),
				ast.NewOrderedList(
					ast.NewListItem(
						ast.NewParagraph(
							ast.NewText("item 1"),
						),
					),
					ast.NewListItem(
						ast.NewParagraph(
							ast.NewText("item 2"),
						),
					),
				),
				ast.NewUnorderedList(
					ast.NewListItem(
						ast.NewStrong(
							ast.NewText("item1"),
						),
					),
					ast.NewListItem(
						ast.NewItalic(
							ast.NewText("item2"),
						),
					),
					ast.NewListItem(
						ast.NewStrikeThrough(
							ast.NewText("item3"),
						),
					),
				),
				ast.NewTable(
					[]ast.TableColumnDefinition{
						{
							Align: ast.TableColumnAlignLeft,
						},
						{
							Align: ast.TableColumnAlignCenter,
						},
						{
							Align: ast.TableColumnAlignRight,
						},
					},
					ast.NewTableRow(
						ast.NewTableCell(
							ast.NewText("head 1"),
						),
						ast.NewTableCell(
							ast.NewText("head 2"),
						),
						ast.NewTableCell(
							ast.NewText("head 3"),
						),
					),
					ast.NewTableRow(
						ast.NewTableCell(
							ast.NewText("body 1-1"),
						),
						ast.NewTableCell(
							ast.NewText("body 1-2"),
						),
						ast.NewTableCell(
							ast.NewText("body 1-3"),
						),
					),
					ast.NewTableRow(
						ast.NewTableCell(
							ast.NewText("body 2-1"),
						),
						ast.NewTableCell(
							ast.NewText("body 2-2"),
						),
						ast.NewTableCell(
							ast.NewText("body 2-3"),
						),
					),
				),
				ast.NewCodeBlock([]string{
					`package main`,
					``,
					`import "fmt"`,
					``,
					`func main() {`,
					`	fmt.Println("Hello, World!")`,
					`}`,
				}, "go"),
				ast.NewBlockQuote(
					ast.NewItalic(ast.NewText("hello")),
					ast.NewStrong(ast.NewText("world!")),
				),
			},
			want: strings.Join([]string{
				`<h1>heading</h1>`,
				`<p><code>paragraph</code>: <a href="https://example.com">link</a>!<br><img src="https://example.com/image.png" alt="awesome image"></p>`,
				`<ol>`,
				`<li><p>item 1</p></li>`,
				`<li><p>item 2</p></li>`,
				`</ol>`,
				`<ul>`,
				`<li><strong>item1</strong></li>`,
				`<li><em>item2</em></li>`,
				`<li><del>item3</del></li>`,
				`</ul>`,
				`<table>`,
				`<tr><td>head 1</td><td>head 2</td><td>head 3</td></tr>`,
				`<tr><td>body 1-1</td><td>body 1-2</td><td>body 1-3</td></tr>`,
				`<tr><td>body 2-1</td><td>body 2-2</td><td>body 2-3</td></tr>`,
				`</table>`,
				`<pre><code class="language-go">`,
				"package main\n",
				"\n",
				"import \"fmt\"\n",
				"\n",
				"func main() {\n",
				"	fmt.Println(\"Hello, World!\")\n",
				"}",
				`</code></pre>`,
				`<blockquote><em>hello</em><strong>world!</strong></blockquote>`,
			}, ""),
		},
	}
	for _, tt := range tests {
		r := NewRenderer()

		t.Run(tt.name, func(t *testing.T) {
			if got := r.Render(tt.nodes); got != tt.want {
				t.Errorf("Render() = \n%v, \nwant \n%v", got, tt.want)
			}
		})
	}
}
