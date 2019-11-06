package main

import (
	"fmt"
	"github.com/pefish/study-compiler-theory/lexer"
)

func main() {
	script := `var a uint64 = 23`
	myLexer := lexer.NewLexer()
	tokens := myLexer.Tokenize(script)
	fmt.Println(tokens)
}
