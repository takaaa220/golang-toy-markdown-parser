package parser

import (
	"reflect"
	"testing"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

var input = []string{
	"# heading",
	"## heading",
	"",
	"- hello",
	"  ### heading",
	"  text",
	"- world",
	"1. list",
	"2. list",
	"",
	"```go",
	`fmt.Print("hello")`,
	"```",
	"",
	"> blockquote1",
	"> blockquote2",
}

func TestParser_block(t *testing.T) {
	type fields struct {
		lines      []string
		lineCursor int
	}
	tests := []struct {
		name          string
		fields        fields
		currentIndent int
		want          ast.Node
		wantErr       bool
	}{
		{
			name: "test1",
			fields: fields{
				lines:      input,
				lineCursor: -1,
			},
			want: ast.HeadingNode(1, ast.TextNode("heading")),
		},
		{
			name: "test2",
			fields: fields{
				lines:      input,
				lineCursor: 0,
			},
			want: ast.HeadingNode(2, ast.TextNode("heading")),
		},
		{
			name: "test3",
			fields: fields{
				lines:      input,
				lineCursor: 2,
			},
			want: ast.UnorderedListNode(
				ast.ListItemNode(
					ast.TextNode("hello"),
					ast.HeadingNode(3, ast.TextNode("heading")),
					ast.ParagraphNode(ast.TextNode("text")),
				),
				ast.ListItemNode(
					ast.TextNode("world"),
				),
			),
		},
		{
			name: "test4",
			fields: fields{
				lines:      input,
				lineCursor: 6,
			},
			want: ast.OrderedListNode(
				ast.ListItemNode(
					ast.TextNode("list"),
				),
				ast.ListItemNode(
					ast.TextNode("list"),
				),
			),
		},
		{
			name: "test5",
			fields: fields{
				lines:      input,
				lineCursor: 9,
			},
			want: ast.CodeBlockNode([]string{
				`fmt.Print("hello")`,
			}, "go"),
		},
		{
			name: "test6",
			fields: fields{
				lines:      input,
				lineCursor: 13,
			},
			want: ast.BlockQuoteNode(
				ast.TextNode("blockquote1"),
				ast.TextNode("blockquote2"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines := []Line{}
			for _, line := range tt.fields.lines {
				lines = append(lines, Line{line})
			}

			p := &Parser{
				lines:      lines,
				lineCursor: tt.fields.lineCursor,
			}
			state := p.newState()
			got, err := p.block(tt.currentIndent, state)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.block() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.block() = \n%v,\n want \n%v", got, tt.want)
			}
		})
	}
}
