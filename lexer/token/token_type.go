package token

type TokenType = byte

const (
	Null  TokenType = iota
	Assignment // =
	Var
	Uint64
	Identifier //标识符

	IntLiteral    //整型字面量
	StringLiteral //字符串字面量
)
