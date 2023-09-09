package parser

import (
	"reflect"
	"testing"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func TestParser_paragraph(t *testing.T) {
	tests := []struct {
		input         string
		currentIndent int
		want          ast.Node
		wantErr       bool
	}{
		{
			input: `hello world`,
			want: ast.ParagraphNode(
				ast.TextNode("hello world"),
			),
		},
		{
			input: "",
			want:  ast.EmptyNode(),
		},
		{
			input: `he**llo** world
YEAH!`,
			want: ast.ParagraphNode(
				ast.TextNode("he"),
				ast.StrongNode(ast.TextNode("llo")),
				ast.TextNode(" world"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			p := NewParser(tt.input)
			got, err := p.paragraph(tt.currentIndent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.paragraph() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.paragraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
