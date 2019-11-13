package lexer

import (
	"fmt"
	"github.com/pefish/study-compiler-theory/lexer/dfa"
	"github.com/pefish/study-compiler-theory/lexer/token"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	Tokens []*token.Token // 存放解析出来的所有token
}

func NewLexer() *Lexer {
	return &Lexer{
		Tokens: []*token.Token{},
	}
}

func (this *Lexer) ToString() string {
	result := ``
	for _, token1 := range this.Tokens {
		result += fmt.Sprintf("%d: %s\n", token1.Type, token1.Text)
	}
	return result
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9' || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
}

// var a uint64 = 23
func (this *Lexer) processToken(ch rune) dfa.DfaState {
	var state dfa.DfaState
	if ch == ' ' || ch == '\t' || ch == '\n' || ch == ';' {
		state = dfa.Milestone
	} else if isDigit(ch) {
		state = dfa.IntLiteral
	} else {
		state = dfa.Identifier
	}
	return state
}

func (this *Lexer) Tokenize(script string) error {
	tempToken := &token.Token{}
	script += " "
	state := dfa.Identifier
	bytes := []byte(script)
	for _, v := range bytes {
		ch := rune(v)
		state = this.processToken(ch)
		switch state {
		case dfa.Identifier:
			tempToken.Text += string(ch)
			tempToken.Type = token.Identifier
		case dfa.IntLiteral:
			tempToken.Text += string(ch)
			tempToken.Type = token.IntLiteral
		case dfa.Milestone:
			if tempToken.Text == `var` {
				tempToken.Type = token.Var
			} else if tempToken.Text == `uint64` {
				tempToken.Type = token.Uint64
			} else if tempToken.Text == `=` {
				tempToken.Type = token.Assignment
			} else if tempToken.Text == `+` {
				tempToken.Type = token.Add
			} else if tempToken.Text == `*` {
				tempToken.Type = token.Multi
			}
			this.Tokens = append(this.Tokens, tempToken)
			tempToken = token.NewToken()
		}
	}
	return nil
}
