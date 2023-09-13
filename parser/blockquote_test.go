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
		currentCursor int
		currentIndent int
		want          ast.Node
		wantErr       bool
		wantCursor    int
		state         blockParsedState
	}{
		{
			input: `> blockquote1
> blockquote2`,
			want: ast.NewBlockQuote(
				ast.NewText("blockquote1"),
				ast.NewText("blockquote2"),
			),
			wantCursor: 1,
		},
		{
			input: `> blockquote1
> blockquote2
hello world,
`,
			want: ast.NewBlockQuote(
				ast.NewText("blockquote1"),
				ast.NewText("blockquote2"),
			),
			wantCursor: 1,
		},
		{
			input: strings.Join([]string{
				" > blockquote1",
				" > blockquote2",
			}, "\n"),
			currentIndent: 1,
			want: ast.NewBlockQuote(
				ast.NewText("blockquote1"),
				ast.NewText("blockquote2"),
			),
			wantCursor: 1,
		},
		{
			input: strings.Join([]string{
				"> block**quote**1",
				"> blockquote2",
			}, "\n"),
			currentIndent: 0,
			want: ast.NewBlockQuote(
				ast.NewText("block"),
				ast.NewStrong(ast.NewText("quote")),
				ast.NewText("1"),
				ast.NewText("blockquote2"),
			),
			wantCursor: 1,
		},
		{
			input: strings.Join([]string{
				"> blockquote1",
				"> blockquote2",
				"hello world",
			}, "\n"),
			currentCursor: 1,
			wantErr:       true,
			wantCursor:    1,
			state: blockParsedState{
				lines: []Line{},
				from:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			p := NewParser(tt.input)
			if tt.currentCursor != 0 {
				p.lineCursor = tt.currentCursor
			}

			state := p.newState()
			got, err := p.blockquote(tt.currentIndent, state)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Parser.blockquote() error = %v, wantErr %v", err, tt.wantErr)
				}

				if !reflect.DeepEqual(tt.state, err.(BlockParseError).State) {
					t.Errorf("Parser.blockquote() state = %v, want %v", err.(BlockParseError).State, tt.state)
				}

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
