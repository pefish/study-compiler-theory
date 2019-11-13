package main

import (
	"fmt"
	"github.com/pefish/study-compiler-theory/ast"
	"github.com/pefish/study-compiler-theory/calc"
	"github.com/pefish/study-compiler-theory/lexer"
	"runtime/debug"
)

// 一个计算器的语法分析（生成抽象语法树）
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
		}
	}()
	script := `var a uint64 = 23 * 2 + 11 * 3`
	myLexer := lexer.NewLexer()
	err := myLexer.Tokenize(script)
	if err != nil {
		panic(err)
	}
	//fmt.Println(myLexer.ToString())
	calculator := calc.Calculator{}
	node := calculator.MustUint64Declare(myLexer.Tokens)
	ast.DumpNode(node, ``)
}
