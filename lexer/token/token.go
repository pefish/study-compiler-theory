package token

type Token struct {
	Type TokenType
	Text string
}

func (this *Token) Empty() {
	this.Type = Null
	this.Text = ``
}
