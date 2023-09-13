package parser

import (
	"reflect"
	"testing"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func TestParser_heading(t *testing.T) {
	tests := []struct {
		input         string
		currentIndent int
		want          ast.Node
		wantErr       bool
		state         blockParsedState
	}{
		{
			input: "# heading",
			want:  ast.HeadingNode(1, ast.TextNode("heading")),
		},
		{
			input: "##  heading",
			want:  ast.HeadingNode(2, ast.TextNode("heading")),
		},
		{
			input:         "  ### heading",
			currentIndent: 2,
			want:          ast.HeadingNode(3, ast.TextNode("heading")),
		},
		{
			input:         "  ### he**ad**ing",
			currentIndent: 2,
			want: ast.HeadingNode(3,
				ast.TextNode("he"),
				ast.StrongNode(ast.TextNode("ad")),
				ast.TextNode("ing"),
			),
		},
		{
			input:   "###heading",
			wantErr: true,
			state: blockParsedState{
				lines: []Line{{"###heading"}},
				from:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			p := NewParser(tt.input)
			state := p.newState()
			got, err := p.heading(tt.currentIndent, state)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Parser.heading() error = %v, wantErr %v", err, tt.wantErr)
				}

				if !reflect.DeepEqual(tt.state, err.(BlockParseError).State) {
					t.Errorf("Parser.heading() state = %v, want %v", err.(BlockParseError).State, tt.state)
				}

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.heading() = %v, want %v", got, tt.want)
			}
			if p.lineCursor != 0 {
				t.Errorf("Parser.heading() lineCursor = %v, want %v", p.lineCursor, 0)
			}
		})
	}
}
