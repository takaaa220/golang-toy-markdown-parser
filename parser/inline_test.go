package parser

import (
	"reflect"
	"testing"

	"github.com/takaaa220/golang-toy-markdown-parser/ast"
)

func Test_inline(t *testing.T) {
	tests := []struct {
		text    string
		want    []ast.Node
		wantErr bool
	}{
		{
			text: "hello",
			want: []ast.Node{
				ast.TextNode("hello"),
			},
		},
		{
			text: "hello `world` yeah",
			want: []ast.Node{
				ast.TextNode("hello "),
				ast.CodeNode("world"),
				ast.TextNode(" yeah"),
			},
		},
		{
			text: "hello [world](https://example.com) yeah",
			want: []ast.Node{
				ast.TextNode("hello "),
				ast.LinkNode("https://example.com", ast.TextNode("world")),
				ast.TextNode(" yeah"),
			},
		},
		{
			text: "hello ![world](https://example.com) yeah",
			want: []ast.Node{
				ast.TextNode("hello "),
				ast.ImageNode("world", "https://example.com"),
				ast.TextNode(" yeah"),
			},
		},
		{
			text: "hello **world** yeah",
			want: []ast.Node{
				ast.TextNode("hello "),
				ast.StrongNode(ast.TextNode("world")),
				ast.TextNode(" yeah"),
			},
		},
		{
			text: "hello *world* yeah",
			want: []ast.Node{
				ast.TextNode("hello "),
				ast.ItalicNode(ast.TextNode("world")),
				ast.TextNode(" yeah"),
			},
		},
		{
			text: "hello ~~world~~ yeah",
			want: []ast.Node{
				ast.TextNode("hello "),
				ast.StrikeThroughNode(ast.TextNode("world")),
				ast.TextNode(" yeah"),
			},
		},
		{
			text: "hello **world yeah",
			want: []ast.Node{
				ast.TextNode("hello **world yeah"),
			},
		},
		{
			text: "hello *world yeah",
			want: []ast.Node{
				ast.TextNode("hello *world yeah"),
			},
		},
		// {
		// 	// TODO: Support this case
		// 	text: "***hello** world* yeah",
		// 	want: []ast.Node{
		// 		ast.ItalicNode(
		// 			ast.StrongNode(
		// 				ast.TextNode("hello"),
		// 			),
		// 			ast.TextNode(" world"),
		// 		),
		// 		ast.TextNode(" yeah"),
		// 	},
		// },
		{
			text: "***hello* world** yeah",
			want: []ast.Node{
				ast.StrongNode(
					ast.ItalicNode(
						ast.TextNode("hello"),
					),
					ast.TextNode(" world"),
				),
				ast.TextNode(" yeah"),
			},
		},
		{
			text: "~hello **world**~ yeah",
			want: []ast.Node{
				ast.StrikeThroughNode(
					ast.TextNode("hello "),
					ast.StrongNode(
						ast.TextNode("world"),
					),
				),
				ast.TextNode(" yeah"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			got, err := inline(tt.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("inline() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inline() = \n%v\n, want \n%v", got, tt.want)
			}
		})
	}
}
