package calc

import (
	"errors"
	"github.com/pefish/study-compiler-theory/ast"
	"github.com/pefish/study-compiler-theory/lexer/token"
)

type Calculator struct {

}


func (this *Calculator) MustUint64Declare(tokens []*token.Token) *ast.AstNode {
	var index uint64 = 0
	if tokens[index] == nil || tokens[index].Type != token.Var {
		panic(errors.New(`语法错误`))
	}
	index++
	if tokens[index] == nil || tokens[index].Type != token.Identifier {
		panic(errors.New(`语法错误`))
	}
	index++
	if tokens[index] == nil || tokens[index].Type != token.Uint64 {
		panic(errors.New(`语法错误`))
	}
	index++
	if tokens[index] == nil || tokens[index].Type != token.Assignment {
		panic(errors.New(`语法错误`))
	}
	index++
	return this.mustAdditive(index, tokens)
}

func (this *Calculator) mustAdditive(index uint64, tokens []*token.Token) *ast.AstNode {
	node, err := this.additive(index, tokens)
	if err != nil {
		panic(err)
	}
	return node
}

// 1 * 2 + 3 * 4 这种表达式
func (this *Calculator) additive(index uint64, tokens []*token.Token) (*ast.AstNode, error) {
	if tokens[index] == nil || tokens[index].Type != token.IntLiteral {
		return nil, errors.New(`语法错误1`)
	}
	child1, index1, err := this.multiplicative(index, tokens)
	if err != nil {
		return nil, err
	}
	index = index1
	if tokens[index] == nil || tokens[index].Type != token.Add {
		return nil, errors.New(`语法错误2`)
	}
	node := &ast.AstNode{
		NodeType: token.Add,
		Text: tokens[index].Text,
	}
	child1.Parent = node
	node.AddChild(child1)
	index++
	if tokens[index] == nil || tokens[index].Type != token.IntLiteral {
		return nil, errors.New(`语法错误3`)
	}
	child2, _, err := this.multiplicative(index, tokens)
	child2.Parent = node
	node.AddChild(child2)
	return node, nil
}

/**
2 * 3 这种基础表达式
 */
func (this *Calculator) multiplicative(index uint64, tokens []*token.Token) (*ast.AstNode, uint64, error) {
	if tokens[index] == nil || tokens[index].Type != token.IntLiteral {
		return nil, 0, errors.New(`语法错误4`)
	}
	child1 := &ast.AstNode{
		NodeType: token.IntLiteral,
		Text: tokens[index].Text,
	}
	index++
	if tokens[index] == nil || tokens[index].Type != token.Multi {
		return child1, index, nil
	}
	node := &ast.AstNode{
		NodeType: token.Multi,
		Text: tokens[index].Text,
	}
	child1.Parent = node
	node.AddChild(child1)
	index++
	if tokens[index] == nil || tokens[index].Type != token.IntLiteral {
		return nil, 0, errors.New(`语法错误6`)
	}
	child2 := &ast.AstNode{
		NodeType: token.IntLiteral,
		Text: tokens[index].Text,
	}
	child2.Parent = node
	node.AddChild(child2)
	index++
	return node, index, nil  // 返回下一个要处理的索引
}

