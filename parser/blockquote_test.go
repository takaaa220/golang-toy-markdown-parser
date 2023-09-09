package parser

import (
	"reflect"
	"strings"
	"testing"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func TestParser_blockquote(t *testing.T) {
	tests := []struct {
		input         string
		currentIndent int
		want          ast.Node
		wantErr       bool
		wantCursor    int
	}{
		{
			input: `> blockquote1
> blockquote2`,
			want: ast.BlockQuoteNode(
				ast.TextNode("blockquote1"),
				ast.TextNode("blockquote2"),
			),
			wantCursor: 1,
		},
		{
			input: `> blockquote1
> blockquote2
hello world,
`,
			want: ast.BlockQuoteNode(
				ast.TextNode("blockquote1"),
				ast.TextNode("blockquote2"),
			),
			wantCursor: 1,
		},
		{
			input: strings.Join([]string{
				" > blockquote1",
				" > blockquote2",
			}, "\n"),
			currentIndent: 1,
			want: ast.BlockQuoteNode(
				ast.TextNode("blockquote1"),
				ast.TextNode("blockquote2"),
			),
			wantCursor: 1,
		},
		{
			input: strings.Join([]string{
				"> block**quote**1",
				"> blockquote2",
			}, "\n"),
			currentIndent: 0,
			want: ast.BlockQuoteNode(
				ast.TextNode("block"),
				ast.StrongNode(ast.TextNode("quote")),
				ast.TextNode("1"),
				ast.TextNode("blockquote2"),
			),
			wantCursor: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			p := NewParser(tt.input)
			got, err := p.blockquote(tt.currentIndent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.blockquote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.blockquote() = %v, want %v", got, tt.want)
			}
			if p.lineCursor != tt.wantCursor {
				t.Errorf("Parser.blockquote() lineCursor = %v, want %v", p.lineCursor, len(p.lines)-1)
			}
		})
	}
}
