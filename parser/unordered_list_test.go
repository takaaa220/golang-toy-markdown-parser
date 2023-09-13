package parser

import (
	"reflect"
	"strings"
	"testing"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func TestParser_unorderedList(t *testing.T) {
	tests := []struct {
		input         string
		currentIndent int
		want          ast.Node
		wantErr       bool
		state         blockParsedState
	}{
		{
			input: strings.Join([]string{
				"- list1",
				" - list1-1",
				"  - list1-1-1",
				" - list1-2",
				"  - list1-2-1",
				"- list2",
			}, "\n"),
			want: ast.NewUnorderedList(
				ast.NewListItem(
					ast.NewText("list1"),
					ast.NewUnorderedList(
						ast.NewListItem(
							ast.NewText("list1-1"),
							ast.NewUnorderedList(
								ast.NewListItem(
									ast.NewText("list1-1-1"),
								),
							),
						),
						ast.NewListItem(
							ast.NewText("list1-2"),
							ast.NewUnorderedList(
								ast.NewListItem(
									ast.NewText("list1-2-1"),
								),
							),
						),
					),
				),
				ast.NewListItem(
					ast.NewText("list2"),
				),
			),
		},
		{
			input: strings.Join([]string{
				"- list1",
				"* list2",
			}, "\n"),
			want: ast.NewUnorderedList(
				ast.NewListItem(
					ast.NewText("list1"),
				),
			),
		},
		{
			input: strings.Join([]string{
				"- list1",
				" # heading",
				" - list1-1",
				" - list1-2",
				"- list2",
			}, "\n"),
			currentIndent: 0,
			want: ast.NewUnorderedList(
				ast.NewListItem(
					ast.NewText("list1"),
					ast.NewHeading(1, ast.NewText("heading")),
					ast.NewUnorderedList(
						ast.NewListItem(
							ast.NewText("list1-1"),
						),
						ast.NewListItem(
							ast.NewText("list1-2"),
						),
					),
				),
				ast.NewListItem(
					ast.NewText("list2"),
				),
			),
		},
		{
			input: strings.Join([]string{
				"- l**is**t1",
				"- list*2*",
			}, "\n"),
			want: ast.NewUnorderedList(
				ast.NewListItem(
					ast.NewText("l"),
					ast.NewStrong(ast.NewText("is")),
					ast.NewText("t1"),
				),
				ast.NewListItem(
					ast.NewText("list"),
					ast.NewItalic(ast.NewText("2")),
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
			state := p.newState()
			got, err := p.unorderedList(tt.currentIndent, state)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Parser.unorderedList() error = %v, wantErr %v", err, tt.wantErr)
				}

				if !reflect.DeepEqual(tt.state, err.(BlockParseError).State) {
					t.Errorf("Parser.unorderedList() state = %v, want %v", err.(BlockParseError).State, tt.state)
				}

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.unorderedList() = \n%v,\n want = \n%v", got, tt.want)
			}
		})
	}
}
