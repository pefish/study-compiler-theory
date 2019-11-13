package main

import (
	"fmt"
	"github.com/pefish/study-compiler-theory/lexer"
)

// 词法分析
func main() {
	script := `var a uint64 = 23`
	myLexer := lexer.NewLexer()
	err := myLexer.Tokenize(script)
	if err != nil {
		panic(err)
	}
	fmt.Println(myLexer.Tokens)
}
