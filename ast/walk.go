package ast

func Walk(node Node, fn func(Node)) {
	fn(node)

	// for _, child := range node.Children() {
	// 	Walk(child, fn)
	// }
}
