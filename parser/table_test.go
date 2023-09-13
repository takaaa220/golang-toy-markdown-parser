package parser

import (
	"reflect"
	"strings"
	"testing"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func TestParser_table(t *testing.T) {
	tests := []struct {
		input         string
		currentIndent int
		want          ast.Node
		wantErr       bool
	}{
		{
			input: strings.Join([]string{
				"| aaa | bbb | ccc | ddd |",
				"| --- | :--- | :---: | ---: |",
				"| 1 | 2 | 3 | 4 |",
				"| a | b | c | d |",
			}, "\n"),
			want: ast.TableNode(
				[]ast.TableColumnDefinition{
					{
						Align: ast.TableColumnAlignLeft,
					},
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
				ast.TableRowNode(
					ast.TableCellNode(
						ast.TextNode("aaa"),
					),
					ast.TableCellNode(
						ast.TextNode("bbb"),
					),
					ast.TableCellNode(
						ast.TextNode("ccc"),
					),
					ast.TableCellNode(
						ast.TextNode("ddd"),
					),
				),
				ast.TableRowNode(
					ast.TableCellNode(
						ast.TextNode("1"),
					),
					ast.TableCellNode(
						ast.TextNode("2"),
					),
					ast.TableCellNode(
						ast.TextNode("3"),
					),
					ast.TableCellNode(
						ast.TextNode("4"),
					),
				),
				ast.TableRowNode(
					ast.TableCellNode(
						ast.TextNode("a"),
					),
					ast.TableCellNode(
						ast.TextNode("b"),
					),
					ast.TableCellNode(
						ast.TextNode("c"),
					),
					ast.TableCellNode(
						ast.TextNode("d"),
					),
				),
			),
		},
		{
			input: strings.Join([]string{
				"| aaa | bbb | ccc |",
				"| --- | :--- | :---: |",
			}, "\n"),
			want:    ast.Node{},
			wantErr: true,
		},
		{
			input: strings.Join([]string{
				"| aaa | bbb | ccc ",
				"| --- | :--- | :---: |",
				"| 1 | 2 | 3 |",
			}, "\n"),
			want:    ast.Node{},
			wantErr: true,
		},
		{
			input: strings.Join([]string{
				"| aaa | bbb | ccc |",
				"| --- | :--- | :---: |",
				"| 1 | 2 |",
			}, "\n"),
			want:    ast.Node{},
			wantErr: true,
		},
		{
			input: strings.Join([]string{
				"| aaa | bbb | ccc |",
				"| -*-- | *--- | ---* |",
				"| 1 | 2 | 3 |",
			}, "\n"),
			want:    ast.Node{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			p := NewParser(tt.input)

			got, err := p.table(tt.currentIndent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.table() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n%v\n\n%v", got, tt.want)
			}
		})
	}
}
