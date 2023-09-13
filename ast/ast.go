package ast

import (
	"fmt"
	"strings"
)

type NodeType string

type Node interface {
	Type() NodeType
	Children() []Node
	AppendChildren(children ...Node)
	Dump() string
}

type NodeBase struct {
	nodeType NodeType
	children []Node
}

func NewNodeBase(nodeType NodeType, children ...Node) *NodeBase {
	return &NodeBase{
		nodeType: nodeType,
		children: children,
	}
}

func (n *NodeBase) Type() NodeType {
	return n.nodeType
}

func (n *NodeBase) Children() []Node {
	return n.children
}

func (n NodeBase) Dump() string {
	return ""
}

func (n *NodeBase) AppendChildren(children ...Node) {
	n.children = append(n.children, children...)
}

func (n *NodeBase) String() string {
	return _innerNodeString(n, 0)
}

func _innerNodeString(n Node, indent int) string {
	text := fmt.Sprintf("%s%s %s", strings.Repeat(" ", indent), n.Type(), n.Dump())

	for _, child := range n.Children() {
		text += "\n" + _innerNodeString(child, indent+2)
	}

	return text
}
