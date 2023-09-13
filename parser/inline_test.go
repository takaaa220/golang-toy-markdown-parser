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
				ast.NewText("hello"),
			},
		},
		{
			text: "hello `world` yeah",
			want: []ast.Node{
				ast.NewText("hello "),
				ast.NewCode("world"),
				ast.NewText(" yeah"),
			},
		},
		{
			text: "hello [world](https://example.com) yeah",
			want: []ast.Node{
				ast.NewText("hello "),
				ast.NewLink("https://example.com", ast.NewText("world")),
				ast.NewText(" yeah"),
			},
		},
		{
			text: "hello ![world](https://example.com) yeah",
			want: []ast.Node{
				ast.NewText("hello "),
				ast.NewImage("world", "https://example.com"),
				ast.NewText(" yeah"),
			},
		},
		{
			text: "hello **world** yeah",
			want: []ast.Node{
				ast.NewText("hello "),
				ast.NewStrong(ast.NewText("world")),
				ast.NewText(" yeah"),
			},
		},
		{
			text: "hello *world* yeah",
			want: []ast.Node{
				ast.NewText("hello "),
				ast.NewItalic(ast.NewText("world")),
				ast.NewText(" yeah"),
			},
		},
		{
			text: "hello ~~world~~ yeah",
			want: []ast.Node{
				ast.NewText("hello "),
				ast.NewStrikeThrough(ast.NewText("world")),
				ast.NewText(" yeah"),
			},
		},
		{
			text: "hello **world yeah",
			want: []ast.Node{
				ast.NewText("hello **world yeah"),
			},
		},
		{
			text: "hello *world yeah",
			want: []ast.Node{
				ast.NewText("hello *world yeah"),
			},
		},
		// {
		// 	// TODO: Support this case
		// 	text: "***hello** world* yeah",
		// 	want: []ast.Node{
		// 		ast.NewItalic(
		// 			ast.NewStrong(
		// 				ast.NewText("hello"),
		// 			),
		// 			ast.NewText(" world"),
		// 		),
		// 		ast.NewText(" yeah"),
		// 	},
		// },
		{
			text: "***hello* world** yeah",
			want: []ast.Node{
				ast.NewStrong(
					ast.NewItalic(
						ast.NewText("hello"),
					),
					ast.NewText(" world"),
				),
				ast.NewText(" yeah"),
			},
		},
		{
			text: "~hello **world**~ yeah",
			want: []ast.Node{
				ast.NewStrikeThrough(
					ast.NewText("hello "),
					ast.NewStrong(
						ast.NewText("world"),
					),
				),
				ast.NewText(" yeah"),
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
