package parser

import (
	"reflect"
	"testing"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func TestParser_unorderedList(t *testing.T) {
	tests := []struct {
		input         string
		currentIndent int
		want          ast.Node
		wantErr       bool
	}{
		{
			input: `- list1
			- list1-1
				- list1-1-1
			- list1-2
				- list1-2-1
- list2`,
			currentIndent: 0,
			want: ast.UnorderedListNode(
				ast.ListItemNode(
					ast.TextNode("list1"),
					ast.UnorderedListNode(
						ast.ListItemNode(
							ast.TextNode("list1-1"),
							ast.UnorderedListNode(
								ast.ListItemNode(
									ast.TextNode("list1-1-1"),
								),
							),
						),
						ast.ListItemNode(
							ast.TextNode("list1-2"),
							ast.UnorderedListNode(
								ast.ListItemNode(
									ast.TextNode("list1-2-1"),
								),
							),
						),
					),
				),
				ast.ListItemNode(
					ast.TextNode("list2"),
				),
			),
			wantErr: false,
		},
		{
			input: string(`- list1
* list2`),
			currentIndent: 0,
			want: ast.UnorderedListNode(ast.ListItemNode(
				ast.TextNode("list1"),
			)),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			l := NewParser(tt.input)
			got, err := l.unorderedList(tt.currentIndent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.unorderedList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.unorderedList() = \n%v,\n want = \n%v", got, tt.want)
			}
		})
	}
}
