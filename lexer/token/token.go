package token

import "fmt"

type Token struct {
	Type TokenType
	Text string
}

func NewToken() *Token {
	tempToken := Token{
		Type: Null,
		Text: ``,
	}
	return &tempToken
}

func (this *Token) ToString() string {
	return fmt.Sprintf(`%d: %s`, this.Type, this.Text)
}
