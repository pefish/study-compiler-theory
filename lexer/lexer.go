package lexer

import (
	"github.com/pefish/study-compiler-theory/lexer/dfa"
	"github.com/pefish/study-compiler-theory/lexer/token"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	Tokens []token.Token // 存放解析出来的所有token
	token  token.Token   // 当前正在解析的Token
}

func NewLexer() *Lexer {
	return &Lexer{
		Tokens: []token.Token{},
		token: token.Token{
			Text: ``,
			Type: token.Null,
		},
	}
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
		state = dfa.Identifier;
	}
	return state
}

// var a uint64 = 23
func (this *Lexer) Tokenize(script string) error {
	script += " "
	state := dfa.Identifier
	bytes := []byte(script)
	for _, v := range bytes {
		ch := rune(v)
		state = this.processToken(ch)
		switch state {
		case dfa.Identifier:
			this.token.Text += string(ch)
			this.token.Type = token.Identifier
		case dfa.IntLiteral:
			this.token.Text += string(ch)
			this.token.Type = token.IntLiteral
		case dfa.Milestone:
			if this.token.Text == `var` {
				this.token.Type = token.Var
			} else if this.token.Text == `uint64` {
				this.token.Type = token.Uint64
			} else if this.token.Text == `=` {
				this.token.Type = token.Assignment
			}
			this.Tokens = append(this.Tokens, this.token)
			this.token.Empty()
		}
	}
	return nil
}
