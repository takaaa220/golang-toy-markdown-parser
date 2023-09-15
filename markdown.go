package golangToyMarkdownParser

import (
	"github.com/takaaa220/golang-toy-markdown-parser/parser"
	"github.com/takaaa220/golang-toy-markdown-parser/renderer"
)

type Options struct {
	Renderer renderer.Renderer
}

func MdToHtml(md string, options Options) (string, error) {
	parser := parser.NewParser(md)

	nodes, err := parser.Parse(0)
	if err != nil {
		return "", err
	}

	r := renderer.NewDefaultRenderer()
	if options.Renderer != nil {
		r = options.Renderer
	}

	return renderer.Render(nodes, r), nil
}
