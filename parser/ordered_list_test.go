package parser

import (
	"reflect"
	"strings"
	"testing"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func TestParser_orderedList(t *testing.T) {
	tests := []struct {
		input         string
		currentIndent int
		want          ast.Node
		wantErr       bool
		state         blockParsedState
	}{
		{
			input: strings.Join([]string{
				"1. list1",
				" 1. list1-1",
				"  1. list1-1-1",
				" 2. list1-2",
				"  1. list1-2-1",
				"2. list2",
			}, "\n"),
			want: ast.OrderedListNode(
				ast.ListItemNode(
					ast.TextNode("list1"),
					ast.OrderedListNode(
						ast.ListItemNode(
							ast.TextNode("list1-1"),
							ast.OrderedListNode(
								ast.ListItemNode(
									ast.TextNode("list1-1-1"),
								),
							),
						),
						ast.ListItemNode(
							ast.TextNode("list1-2"),
							ast.OrderedListNode(
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
		},
		{
			input: strings.Join([]string{
				"1. list1",
				"1. list2",
			}, "\n"),
			want: ast.OrderedListNode(
				ast.ListItemNode(
					ast.TextNode("list1"),
				),
			),
		},
		{
			input:   strings.Join([]string{"hello world"}, "\n"),
			wantErr: true,
			state: blockParsedState{
				lines: []Line{},
				from:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			p := NewParser(tt.input)
			got, err := p.orderedList(tt.currentIndent)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Parser.orderedList() error = %v, wantErr %v", err, tt.wantErr)
				}

				if !reflect.DeepEqual(tt.state, err.(BlockParseError).State) {
					t.Errorf("Parser.orderedList() state = %v, want %v", err.(BlockParseError).State, tt.state)
				}

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.orderedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
