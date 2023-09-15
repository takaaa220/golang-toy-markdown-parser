package golangToyMarkdownParser

import (
	"fmt"
	"testing"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
	"github.com/takaaa220/golang-toy-markdown-parser/renderer"
)

type CustomRenderer struct {
	renderer.DefaultRenderer
}

func (r CustomRenderer) Paragraph(p ast.Paragraph) string {
	return fmt.Sprintf("<p class=\"custom\">%s</p>", renderer.Render(p.Children(), r))
}

func TestMdToHtml(t *testing.T) {
	type args struct {
		md      string
		options Options
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "default renderer",
			args: args{
				md: "# heading\nparagraph\n",
			},
			want:    "<h1>heading</h1><p>paragraph</p>",
			wantErr: false,
		},
		{
			name: "custom renderer",
			args: args{
				md: "# heading\nparagraph\n",
				options: Options{
					Renderer: &CustomRenderer{},
				},
			},
			want:    "<h1>heading</h1><p class=\"custom\">paragraph</p>",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MdToHtml(tt.args.md, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("MdToHtml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MdToHtml() = %v, want %v", got, tt.want)
			}
		})
	}
}
