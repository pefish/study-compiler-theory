package dfa

// Deterministic Finite Automaton 确定性有限状态机
type DfaState = byte

const (
	Initial DfaState = iota
	Identifier  // 标识符状态(词)
	Milestone // 碰到空格、回车、tab、;时
	//GT
	//GE
	//Plus
	//Minus
	//Star
	//Slash
	//SemiColon
	//LeftParen
	//RightParen
	IntLiteral
)
