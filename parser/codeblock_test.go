package parser

import (
	"reflect"
	"strings"
	"testing"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func TestParser_codeblock(t *testing.T) {
	tests := []struct {
		input         string
		currentCursor int
		want          ast.Node
		wantErr       bool
		wantCursor    int
		state         blockParsedState
	}{
		{
			input: strings.Join([]string{
				"```go",
				"v, err := main()",
				"if err != nil {",
				"  return err",
				"}",
				"```",
			}, "\n"),
			want: ast.CodeBlockNode(
				[]string{
					"v, err := main()",
					"if err != nil {",
					"  return err",
					"}",
				},
				"go",
			),
			wantCursor: 5,
		},
		{
			input: strings.Join([]string{
				"  ```go",
				"  v, err := main()",
				"  if err != nil {",
				"    return err",
				"  }",
				"  ```",
			}, "\n"),
			currentCursor: 2,
			want: ast.CodeBlockNode(
				[]string{
					"v, err := main()",
					"if err != nil {",
					"  return err",
					"}",
				},
				"go",
			),
			wantCursor: 5,
		},
		{
			input: strings.Join([]string{
				"```",
				"v, err := main()",
				"if err != nil {",
				"  return err",
				"}",
				"```",
			}, "\n"),
			want: ast.CodeBlockNode(
				[]string{
					"v, err := main()",
					"if err != nil {",
					"  return err",
					"}",
				},
				"",
			),
			wantCursor: 5,
		},
		{
			input: strings.Join([]string{
				"hello world",
			}, "\n"),
			wantErr: true,
			state: blockParsedState{
				lines: []Line{{"hello world"}},
				from:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			p := NewParser(tt.input)
			state := p.newState()

			got, err := p.codeblock(tt.currentCursor, state)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Parser.codeblock() error = %v, wantErr %v", err, tt.wantErr)
				}

				if !reflect.DeepEqual(tt.state, err.(BlockParseError).State) {
					t.Errorf("Parser.codeblock() state = %v, want %v", err.(BlockParseError).State, tt.state)
				}

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.codeblock() = \n%v,\n want \n%v", got, tt.want)
			}
		})
	}
}
