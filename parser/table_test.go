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
		state         blockParsedState
	}{
		{
			input: strings.Join([]string{
				"| aaa | bbb | ccc | ddd |",
				"| --- | :--- | :---: | ---: |",
				"| 1 | 2 | 3 | 4 |",
				"| a | b | c | d |",
			}, "\n"),
			want: ast.NewTable(
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
				ast.NewTableRow(
					ast.NewTableCell(
						ast.NewText("aaa"),
					),
					ast.NewTableCell(
						ast.NewText("bbb"),
					),
					ast.NewTableCell(
						ast.NewText("ccc"),
					),
					ast.NewTableCell(
						ast.NewText("ddd"),
					),
				),
				ast.NewTableRow(
					ast.NewTableCell(
						ast.NewText("1"),
					),
					ast.NewTableCell(
						ast.NewText("2"),
					),
					ast.NewTableCell(
						ast.NewText("3"),
					),
					ast.NewTableCell(
						ast.NewText("4"),
					),
				),
				ast.NewTableRow(
					ast.NewTableCell(
						ast.NewText("a"),
					),
					ast.NewTableCell(
						ast.NewText("b"),
					),
					ast.NewTableCell(
						ast.NewText("c"),
					),
					ast.NewTableCell(
						ast.NewText("d"),
					),
				),
			),
		},
		{
			input: strings.Join([]string{
				"| aaa | bbb | ccc |",
				"| --- | :--- | :---: |",
				"| 1 | 2 |",
			}, "\n"),
			want: ast.NewTable(
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
				},
				ast.NewTableRow(
					ast.NewTableCell(
						ast.NewText("aaa"),
					),
					ast.NewTableCell(
						ast.NewText("bbb"),
					),
					ast.NewTableCell(
						ast.NewText("ccc"),
					),
				),
				ast.NewTableRow(
					ast.NewTableCell(
						ast.NewText("1"),
					),
					ast.NewTableCell(
						ast.NewText("2"),
					),
					ast.NewTableCell(
						ast.NewEmpty(),
					),
				),
			),
		},
		{
			input: strings.Join([]string{
				"| aaa | bbb | ccc |",
				"| --- | :--- | :---: |",
			}, "\n"),
			want:    &ast.NodeBase{},
			wantErr: true,
			state: blockParsedState{
				lines: []Line{
					{"| aaa | bbb | ccc |"},
					{"| --- | :--- | :---: |"},
				},
				from: 0,
			},
		},
		{
			input: strings.Join([]string{
				"| aaa | bbb | ccc ",
				"| --- | :--- | :---: |",
				"| 1 | 2 | 3 |",
			}, "\n"),
			want:    &ast.NodeBase{},
			wantErr: true,
			state: blockParsedState{
				lines: []Line{
					{"| aaa | bbb | ccc "},
					{"| --- | :--- | :---: |"},
				},
				from: 0,
			},
		},
		{
			input: strings.Join([]string{
				"| aaa | bbb | ccc |",
				"| -*-- | *--- | ---* |",
				"| 1 | 2 | 3 |",
			}, "\n"),
			want:    &ast.NodeBase{},
			wantErr: true,
			state: blockParsedState{
				lines: []Line{
					{"| aaa | bbb | ccc |"},
					{"| -*-- | *--- | ---* |"},
				},
				from: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			p := NewParser(tt.input)
			state := p.newState()

			got, err := p.table(tt.currentIndent, state)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Parser.table() error = %v, wantErr %v", err, tt.wantErr)
				}

				if !reflect.DeepEqual(tt.state, err.(BlockParseError).State) {
					t.Errorf("Parser.table() state = %v, want %v", err.(BlockParseError).State, tt.state)
				}

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n%v\n\n%v", got, tt.want)
			}
		})
	}
}
