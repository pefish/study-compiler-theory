package ast

import "fmt"

type AstNode struct {
	Parent   *AstNode
	Children []*AstNode
	NodeType NodeType
	Text     string
}

func (this *AstNode) AddChild(node *AstNode) {
	this.Children = append(this.Children, node)
}

func DumpNode(node *AstNode, indent string) {
	fmt.Printf("%s %s\n", indent, node.Text)
	for _, child := range node.Children {
		DumpNode(child, indent + "\t")
	}
}
