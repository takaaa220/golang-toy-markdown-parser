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
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			p := NewParser(tt.input)
			got, err := p.codeblock(tt.currentCursor)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.codeblock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.codeblock() = \n%v,\n want \n%v", got, tt.want)
			}
		})
	}
}
