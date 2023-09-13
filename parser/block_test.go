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
			want: ast.NewHeading(1, ast.NewText("heading")),
		},
		{
			name: "test2",
			fields: fields{
				lines:      input,
				lineCursor: 0,
			},
			want: ast.NewHeading(2, ast.NewText("heading")),
		},
		{
			name: "test3",
			fields: fields{
				lines:      input,
				lineCursor: 2,
			},
			want: ast.NewUnorderedList(
				ast.NewListItem(
					ast.NewText("hello"),
					ast.NewHeading(3, ast.NewText("heading")),
					ast.NewParagraph(ast.NewText("text")),
				),
				ast.NewListItem(
					ast.NewText("world"),
				),
			),
		},
		{
			name: "test4",
			fields: fields{
				lines:      input,
				lineCursor: 6,
			},
			want: ast.NewOrderedList(
				ast.NewListItem(
					ast.NewText("list"),
				),
				ast.NewListItem(
					ast.NewText("list"),
				),
			),
		},
		{
			name: "test5",
			fields: fields{
				lines:      input,
				lineCursor: 9,
			},
			want: ast.NewCodeBlock([]string{
				`fmt.Print("hello")`,
			}, "go"),
		},
		{
			name: "test6",
			fields: fields{
				lines:      input,
				lineCursor: 13,
			},
			want: ast.NewBlockQuote(
				ast.NewText("blockquote1"),
				ast.NewText("blockquote2"),
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
